package booking

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"ivmanto.com/backend/internal/gcal"
)

// Handler for booking-related HTTP requests.
type Handler struct {
	gcalSvc *gcal.Service
}

// NewHandler creates a new booking handler.
func NewHandler(gcalSvc *gcal.Service) *Handler {
	return &Handler{gcalSvc: gcalSvc}
}

// RegisterRoutes registers the booking routes with a mux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/booking/availability", h.handleGetAvailability)
	mux.HandleFunc("/api/booking/book", h.handleBook)
}

// handleGetAvailability handles GET /api/booking/availability
func (h *Handler) handleGetAvailability(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	if dateStr == "" {
		http.Error(w, "Missing 'date' query parameter", http.StatusBadRequest)
		return
	}

	// The frontend sends a date like "2025-09-01". We must interpret this
	// date in the calendar's local timezone to define the correct 24-hour
	// window for that day. We assume the gcal.Service can provide its
	// timezone location, which it should fetch upon initialization.
	loc := h.gcalSvc.Location()

	// Parse the date string using the calendar's timezone. This creates a
	// time.Time object representing midnight at the beginning of the requested
	// day in that specific timezone.
	day, err := time.ParseInLocation(time.DateOnly, dateStr, loc)
	if err != nil {
		http.Error(w, "Invalid date format, please use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	events, err := h.gcalSvc.GetAvailability(day)
	if err != nil {
		log.Printf("ERROR: getting availability: %v", err)
		http.Error(w, "Could not retrieve availability", http.StatusInternalServerError)
		return
	}

	// We only need to return the start and end times to the frontend.
	type Slot struct {
		Start string `json:"start"`
		End   string `json:"end"`
	}
	var slots []Slot
	for _, item := range events {
		slots = append(slots, Slot{
			Start: item.Start.DateTime,
			End:   item.End.DateTime,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(slots)
}

// handleBook handles POST /api/booking/book
func (h *Handler) handleBook(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Start string `json:"start"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Notes string `json:"notes"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	startTime, err := time.Parse(time.RFC3339, req.Start)
	if err != nil {
		http.Error(w, "Invalid start time format, must be RFC3339", http.StatusBadRequest)
		return
	}

	details := gcal.BookingDetails{
		StartTime: startTime,
		Name:      req.Name,
		Email:     req.Email,
		Notes:     req.Notes,
	}

	updatedEvent, err := h.gcalSvc.BookSlot(details)
	if err != nil {
		if errors.Is(err, gcal.ErrSlotNotFound) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		log.Printf("ERROR: booking slot: %v", err)
		http.Error(w, "Could not book consultation", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(updatedEvent)
}
