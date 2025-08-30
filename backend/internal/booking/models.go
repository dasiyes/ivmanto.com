package booking

import "time"

// TimeSlot represents an available time for a booking.
type TimeSlot struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

// Booking represents a confirmed appointment.
type Booking struct {
	ID        string    `json:"id"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Notes     string    `json:"notes"`
}

// BookingRequest is the payload for creating a new booking.
type BookingRequest struct {
	StartTime time.Time `json:"startTime" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Notes     string    `json:"notes"`
}
