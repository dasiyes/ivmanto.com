package contact

import (
	"encoding/json"
	"log"
	"net/http"

	"ivmanto.com/backend/internal/email"
)

// Handler holds dependencies for the contact handlers.
type Handler struct {
	emailer email.Service
}

// NewHandler creates a new contact handler.
func NewHandler(emailer email.Service) *Handler {
	return &Handler{emailer: emailer}
}

// RegisterRoutes registers the contact routes with a mux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/contact", h.handleContactSubmit)
}

func (h *Handler) handleContactSubmit(w http.ResponseWriter, r *http.Request) {
	// Explicitly check for the POST method. This is a more robust pattern.
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

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
		log.Printf("Failed to send contact email: %v", err)
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message sent successfully"))
}
