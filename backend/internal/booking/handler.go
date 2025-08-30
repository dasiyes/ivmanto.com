package booking

import (
	"encoding/json"
	"net/http"
	"time"
)

// Handler holds dependencies for the booking handlers.
type Handler struct {
	service *BookingService
}

// NewHandler creates a new booking handler.
func NewHandler(service *BookingService) *Handler {
	return &Handler{service: service}
}

// RegisterRoutes registers the booking routes with a mux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/booking/availability", h.handleGetAvailability)
	mux.HandleFunc("/api/booking/book", h.handleCreateBooking)
}

func (h *Handler) handleGetAvailability(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	dateStr := r.URL.Query().Get("date")
	if dateStr == "" {
		http.Error(w, "date query parameter is required", http.StatusBadRequest)
		return
	}

	// Expecting date in YYYY-MM-DD format
	day, err := time.Parse(time.DateOnly, dateStr)
	if err != nil {
		http.Error(w, "invalid date format, use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	slots, err := h.service.GetAvailability(day)
	if err != nil {
		http.Error(w, "failed to get availability", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(slots)
}

func (h *Handler) handleCreateBooking(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req BookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: Add validation for the request struct.
	// Using a library like go-playground/validator is recommended.

	booking, err := h.service.CreateBooking(req)
	if err != nil {
		// This could be a conflict (slot already booked) or other server error.
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(booking)
}
