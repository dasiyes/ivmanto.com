package booking

import (
	"errors"
	"sync"
	"time"
)

// Storage defines the interface for booking persistence.
type Storage interface {
	GetBookingsForDay(day time.Time) ([]Booking, error)
	CreateBooking(booking Booking) error
}

// InMemoryBookingStore is a simple in-memory storage for bookings.
// NOTE: This is for demonstration purposes. In a real application,
// you would replace this with a persistent database like Firestore or Cloud SQL.
type InMemoryBookingStore struct {
	mu       sync.RWMutex
	bookings []Booking
}

// NewInMemoryBookingStore creates a new in-memory booking store.
func NewInMemoryBookingStore() *InMemoryBookingStore {
	return &InMemoryBookingStore{
		bookings: make([]Booking, 0),
	}
}

// GetBookingsForDay returns all bookings for a given day.
func (s *InMemoryBookingStore) GetBookingsForDay(day time.Time) ([]Booking, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var dayBookings []Booking
	year, month, dayOfMonth := day.Date()

	for _, b := range s.bookings {
		bYear, bMonth, bDayOfMonth := b.StartTime.Date()
		if bYear == year && bMonth == month && bDayOfMonth == dayOfMonth {
			dayBookings = append(dayBookings, b)
		}
	}

	return dayBookings, nil
}

// CreateBooking adds a new booking to the store.
func (s *InMemoryBookingStore) CreateBooking(booking Booking) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Simple check for overlapping bookings
	for _, b := range s.bookings {
		if booking.StartTime.Before(b.EndTime) && booking.EndTime.After(b.StartTime) {
			return errors.New("time slot is already booked")
		}
	}

	s.bookings = append(s.bookings, booking)
	return nil
}
