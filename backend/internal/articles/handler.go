package articles

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"sync"
)

// Handler manages article-related HTTP requests, such as likes.
type Handler struct {
	logger *slog.Logger

	// For demonstration, we use a simple in-memory map to store likes.
	// This is NOT suitable for production as data will be lost on restart.
	// In a real-world application, this should be replaced with a persistent
	// data store like Firestore, Redis, or a SQL database.
	likes      map[string]int
	likesMutex sync.RWMutex
}

// NewHandler creates a new articles handler.
func NewHandler(logger *slog.Logger) *Handler {
	return &Handler{
		logger: logger,
		likes:  make(map[string]int),
		// RWMutex is initialized with its zero value.
	}
}

// RegisterRoutes sets up the routing for articles endpoints.
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/articles/{slug}/likes", h.handleGetLikes)
	mux.HandleFunc("POST /api/articles/{slug}/like", h.handleIncrementLike)
	mux.HandleFunc("DELETE /api/articles/{slug}/like", h.handleDecrementLike)
}

// handleGetLikes retrieves the current like count for a specific article.
func (h *Handler) handleGetLikes(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		h.respondError(w, http.StatusBadRequest, "Article slug is required")
		return
	}

	h.likesMutex.RLock()
	// Reading from the map is safe. If the slug doesn't exist, it returns the zero value (0).
	likes := h.likes[slug]
	h.likesMutex.RUnlock()

	response := map[string]interface{}{
		"slug":  slug,
		"likes": likes,
	}

	h.respondJSON(w, http.StatusOK, response)
}

// handleIncrementLike increments the like count for a specific article.
func (h *Handler) handleIncrementLike(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		h.respondError(w, http.StatusBadRequest, "Article slug is required")
		return
	}

	h.likesMutex.Lock()
	h.likes[slug]++
	newLikes := h.likes[slug]
	h.likesMutex.Unlock()

	h.logger.Info("Like incremented", "slug", slug, "new_count", newLikes)

	response := map[string]interface{}{
		"likes": newLikes,
	}

	h.respondJSON(w, http.StatusOK, response)
}

// handleDecrementLike decrements the like count for a specific article.
func (h *Handler) handleDecrementLike(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	if slug == "" {
		h.respondError(w, http.StatusBadRequest, "Article slug is required")
		return
	}

	h.likesMutex.Lock()
	// Prevent the count from going below zero.
	if h.likes[slug] > 0 {
		h.likes[slug]--
	}
	newLikes := h.likes[slug]
	h.likesMutex.Unlock()

	h.logger.Info("Like decremented", "slug", slug, "new_count", newLikes)

	response := map[string]interface{}{
		"likes": newLikes,
	}

	h.respondJSON(w, http.StatusOK, response)
}

// respondJSON is a helper to write a JSON response.
func (h *Handler) respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload != nil {
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			h.logger.Error("could not write JSON response", "error", err)
		}
	}
}

// respondError is a helper to write a JSON error message.
func (h *Handler) respondError(w http.ResponseWriter, status int, message string) {
	h.respondJSON(w, status, map[string]string{"message": message})
}
