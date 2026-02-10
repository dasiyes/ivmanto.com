package blog

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// Handler serves blog article endpoints.
type Handler struct {
	logger *slog.Logger
	cache  *Cache
}

// NewHandler creates a new blog handler.
func NewHandler(logger *slog.Logger, cache *Cache) *Handler {
	return &Handler{
		logger: logger,
		cache:  cache,
	}
}

// RegisterRoutes sets up the routing for blog endpoints.
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/articles", h.handleListArticles)
	mux.HandleFunc("GET /api/articles/{slug}", h.handleGetArticle)
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
