package blog

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"
)

// Handler serves blog article endpoints.
type Handler struct {
	logger    *slog.Logger
	cache     *Cache
	pushToken string // optional shared secret for Pub/Sub push validation
}

// NewHandler creates a new blog handler.
func NewHandler(logger *slog.Logger, cache *Cache, pushToken string) *Handler {
	return &Handler{
		logger:    logger,
		cache:     cache,
		pushToken: pushToken,
	}
}

// RegisterRoutes sets up the routing for blog endpoints.
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/articles", h.handleListArticles)
	mux.HandleFunc("GET /api/articles/{slug}", h.handleGetArticle)
	mux.HandleFunc("POST /api/_internal/pubsub/blog-refresh", h.handlePubSubPush)
}

// handleListArticles returns metadata for all published articles.
func (h *Handler) handleListArticles(w http.ResponseWriter, r *http.Request) {
	articles := h.cache.GetAllPublished()
	h.respondJSON(w, http.StatusOK, articles)
}

// handleGetArticle returns a single article's metadata and HTML content.
func (h *Handler) handleGetArticle(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		h.respondError(w, http.StatusBadRequest, "Article slug is required")
		return
	}

	article := h.cache.GetBySlug(slug)
	if article == nil {
		h.respondError(w, http.StatusNotFound, "Article not found")
		return
	}

	h.respondJSON(w, http.StatusOK, article)
}

// pubSubMessage represents the Pub/Sub push message envelope.
type pubSubMessage struct {
	Message struct {
		Attributes map[string]string `json:"attributes"`
		MessageID  string            `json:"messageId"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

// handlePubSubPush handles GCS notification push messages from Pub/Sub.
// It validates the shared token, filters non-.md files, and schedules a
// debounced cache refresh.
func (h *Handler) handlePubSubPush(w http.ResponseWriter, r *http.Request) {
	// 1. Validate the push token if configured.
	if h.pushToken != "" {
		token := r.URL.Query().Get("token")
		if token != h.pushToken {
			h.logger.Warn("pubsub push rejected: invalid token")
			w.WriteHeader(http.StatusForbidden)
			return
		}
	}

	// 2. Decode the Pub/Sub message.
	var msg pubSubMessage
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		h.logger.Error("failed to decode pubsub message", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 3. Only refresh for .md file changes (ignore metadata.json writes, etc).
	objectID := msg.Message.Attributes["objectId"]
	eventType := msg.Message.Attributes["eventType"]

	if !strings.HasSuffix(objectID, ".md") {
		h.logger.Debug("pubsub notification ignored: not a .md file",
			"object", objectID, "event", eventType)
		w.WriteHeader(http.StatusOK)
		return
	}

	h.logger.Info("pubsub notification: triggering cache refresh",
		"object", objectID, "event", eventType, "messageId", msg.Message.MessageID)

	// 4. Schedule debounced refresh (returns immediately).
	h.cache.Refresh()

	// 5. Acknowledge (200 OK). The actual refresh runs asynchronously.
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload != nil {
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			h.logger.Error("could not write JSON response", "error", err)
		}
	}
}

func (h *Handler) respondError(w http.ResponseWriter, status int, message string) {
	h.respondJSON(w, status, map[string]string{"message": message})
}
