package blog

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log/slog"
	"net/http"
	"strings"
	"time"
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
	mux.HandleFunc("GET /api/_internal/articles/status", h.handleArticlesStatus)
	mux.HandleFunc("GET /api/sitemap-blog.xml", h.handleBlogSitemap)
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

// handleArticlesStatus returns diagnostic information about all articles
// including skipped ones. Protected by the push token.
func (h *Handler) handleArticlesStatus(w http.ResponseWriter, r *http.Request) {
	if h.pushToken != "" {
		token := r.URL.Query().Get("token")
		if token != h.pushToken {
			h.logger.Warn("articles status rejected: invalid token")
			w.WriteHeader(http.StatusForbidden)
			return
		}
	}

	status := h.cache.GetCacheStatus()
	h.respondJSON(w, http.StatusOK, status)
}

// sitemapURLSet is the root element for a sitemap XML document.
type sitemapURLSet struct {
	XMLName xml.Name     `xml:"urlset"`
	XMLNS   string       `xml:"xmlns,attr"`
	URLs    []sitemapURL `xml:"url"`
}

type sitemapURL struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
}

// handleBlogSitemap generates a dynamic sitemap XML for all published articles.
func (h *Handler) handleBlogSitemap(w http.ResponseWriter, r *http.Request) {
	articles := h.cache.GetAllPublished()

	urls := make([]sitemapURL, 0, len(articles)+1)

	// Add the blog index page.
	urls = append(urls, sitemapURL{
		Loc:     "https://ivmanto.com/blog",
		LastMod: time.Now().UTC().Format("2006-01-02"),
	})

	for _, a := range articles {
		urls = append(urls, sitemapURL{
			Loc:     fmt.Sprintf("https://ivmanto.com/blog/%s", a.Slug),
			LastMod: a.Date,
		})
	}

	urlSet := sitemapURLSet{
		XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  urls,
	}

	w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, xml.Header)
	if err := xml.NewEncoder(w).Encode(urlSet); err != nil {
		h.logger.Error("failed to write blog sitemap", "error", err)
	}
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
