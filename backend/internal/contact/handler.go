package contact

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"ivmanto.com/backend/internal/email"
)

// Handler holds dependencies for the contact handlers.
type Handler struct {
	logger  *slog.Logger
	emailer email.Service
}

// NewHandler creates a new contact handler.
func NewHandler(logger *slog.Logger, emailer email.Service) *Handler {
	return &Handler{logger: logger, emailer: emailer}
}

// RegisterRoutes registers the contact routes with a mux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/contact", h.handleContactSubmit)
}

func (h *Handler) handleContactSubmit(w http.ResponseWriter, r *http.Request) {
	// The method is now checked by the mux pattern "POST /api/contact"
	// for Go 1.22+

	var msg email.ContactMessage
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Basic validation
	if msg.Name == "" || msg.Email == "" || msg.Message == "" {
		http.Error(w, "Name, email, and message are required", http.StatusBadRequest)
		return
	}
	// A more robust validation would use a library like go-playground/validator

	if err := h.emailer.SendContactMessage(msg); err != nil {
		// In a real app, you'd log the internal error but not expose details to the client.
		h.logger.Error("Failed to send contact email", "error", err)
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message sent successfully"))
}
