package booking

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"ivmanto.com/backend/internal/email"
	"ivmanto.com/backend/internal/gcal"
)

// Handler manages booking-related HTTP requests.
type Handler struct {
	logger   *slog.Logger
	gcalSvc  gcal.Service
	emailSvc email.Service
}

// NewHandler creates a new booking handler.
func NewHandler(logger *slog.Logger, gcalSvc gcal.Service, emailSvc email.Service) *Handler {
	return &Handler{
		logger:   logger,
		gcalSvc:  gcalSvc,
		emailSvc: emailSvc,
	}
}

// RegisterRoutes sets up the routing for booking endpoints.
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/booking/book", h.handleCreateBooking)
	mux.HandleFunc("GET /api/booking/availability", h.handleGetAvailability)
	mux.HandleFunc("POST /api/booking/cancel", h.handleCancelBooking)
}

type cancelRequest struct {
	Token string `json:"token"`
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

// handleCancelBooking processes a request to cancel a booking.
func (h *Handler) handleCancelBooking(w http.ResponseWriter, r *http.Request) {
	var req cancelRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Token == "" {
		h.respondError(w, http.StatusBadRequest, "Cancellation token is required")
		return
	}

	h.logger.Info("Received cancellation request", "token_prefix", req.Token[:8])

	originalEvent, err := h.gcalSvc.CancelBooking(r.Context(), req.Token)
	if err != nil {
		h.logger.Error("Failed to cancel booking", "token_prefix", req.Token[:8], "error", err)
		if errors.Is(err, gcal.ErrSlotNotFound) {
			h.respondError(w, http.StatusNotFound, "Booking not found. The link may be invalid or expired.")
			return
		}
		h.respondError(w, http.StatusInternalServerError, "An internal error occurred while cancelling the booking.")
		return
	}

	h.logger.Info("Booking cancelled successfully", "event_id", originalEvent.Id)

	// Extract details for notifications
	var clientName, clientEmail string
	if len(originalEvent.Attendees) > 0 {
		clientName = originalEvent.Attendees[0].DisplayName
		clientEmail = originalEvent.Attendees[0].Email
	} else {
		h.logger.Warn("Could not find attendee details on cancelled event. Notifications may be incomplete.", "event_id", originalEvent.Id)
		clientName = "Client" // Fallback
	}
	startTime, _ := time.Parse(time.RFC3339, originalEvent.Start.DateTime)

	// Send notifications. We can run these in goroutines for speed.
	go func() {
		err := h.emailSvc.SendBookingCancellationToClient(clientName, clientEmail, startTime)
		if err != nil {
			h.logger.Error("Failed to send cancellation email to client", "client_email", clientEmail, "error", err)
		}
	}()

	go func() {
		err := h.emailSvc.SendBookingCancellationToAdmin(clientName, clientEmail, startTime)
		if err != nil {
			h.logger.Error("Failed to send cancellation notification to admin", "error", err)
		}
	}()

	h.respondJSON(w, http.StatusOK, map[string]string{"message": "Booking cancelled successfully"})
}

// handleGetAvailability handles requests for available time slots.
func (h *Handler) handleGetAvailability(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("date")
	if dateStr == "" {
		h.respondError(w, http.StatusBadRequest, "date query parameter is required")
		return
	}

	// The frontend sends date in YYYY-MM-DD format.
	// We parse it in the calendar's location to correctly handle timezones.
	day, err := time.ParseInLocation("2006-01-02", dateStr, h.gcalSvc.Location())
	if err != nil {
		h.respondError(w, http.StatusBadRequest, "invalid date format, use YYYY-MM-DD")
		return
	}

	events, err := h.gcalSvc.GetAvailability(day)
	if err != nil {
		h.logger.Error("Failed to get availability from Google Calendar", "error", err)
		h.respondError(w, http.StatusInternalServerError, "Failed to get availability")
		return
	}

	// The frontend expects a specific format.
	type availabilityResponse struct {
		Start string `json:"start"`
		ID    string `json:"id"`
		End   string `json:"end"`
	}

	responseSlots := make([]availabilityResponse, len(events))
	for i, event := range events {
		responseSlots[i] = availabilityResponse{
			Start: event.Start.DateTime,
			ID:    event.Id,
			End:   event.End.DateTime,
		}
	}

	h.respondJSON(w, http.StatusOK, responseSlots)
}

type createBookingRequest struct {
	EventID string `json:"eventId"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Notes   string `json:"notes"`
}

// handleCreateBooking handles a new booking request.
func (h *Handler) handleCreateBooking(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Received POST /api/booking/book request")
	var req createBookingRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Basic validation
	if req.Name == "" || req.Email == "" || req.EventID == "" {
		h.respondError(w, http.StatusBadRequest, "Bad Request: Name, email, and eventId are required")
		return
	}

	bookingDetails := gcal.BookingDetails{
		EventID: req.EventID,
		Name:    req.Name,
		Email:   req.Email,
		Notes:   req.Notes,
	}

	event, err := h.gcalSvc.BookSlot(bookingDetails)
	if err != nil {
		h.logger.Error("BookSlot service call failed", "error", err)
		if errors.Is(err, gcal.ErrSlotNotFound) {
			h.respondError(w, http.StatusConflict, "This time slot is no longer available. Please select another time.")
			return
		}
		h.logger.Error("Failed to create booking in Google Calendar", "error", err)
		h.respondError(w, http.StatusInternalServerError, "An internal error occurred while creating the booking.")
		return
	}

	h.logger.Info("Booking created successfully", "event_id", event.Id)

	// Send confirmation emails in the background
	go func() {
		if err := h.emailSvc.SendBookingConfirmation(req.Name, req.Email, event); err != nil {
			h.logger.Error("Failed to send booking confirmation to client", "client_email", req.Email, "error", err)
		}
	}()
	go func() {
		// We need to parse the start time from the event for the admin notification
		startTime, err := time.Parse(time.RFC3339, event.Start.DateTime)
		if err != nil {
			h.logger.Error("Could not parse start time from booked event for admin email", "error", err)
			return
		}
		if err := h.emailSvc.SendBookingNotificationToAdmin(req.Name, req.Email, startTime, req.Notes); err != nil {
			h.logger.Error("Failed to send booking notification to admin", "error", err)
		}
	}()

	h.respondJSON(w, http.StatusCreated, event)
}
