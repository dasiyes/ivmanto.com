package booking

import (
	"time"

	"github.com/google/uuid"
	"ivmanto.com/backend/internal/config"
	"ivmanto.com/backend/internal/email"
)

// BookingService handles the business logic for bookings.
type BookingService struct {
	store   Storage
	emailer email.Service
	cfg     *config.BookingConfig
}

// NewBookingService creates a new booking service.
func NewBookingService(store Storage, emailer email.Service, cfg *config.BookingConfig) *BookingService {
	return &BookingService{
		store:   store,
		emailer: emailer,
		cfg:     cfg,
	}
}

// GetAvailability returns available time slots for a given day.
func (s *BookingService) GetAvailability(day time.Time) ([]TimeSlot, error) {
	// For simplicity, we'll only allow booking for the next 30 days.
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	if day.Before(today) || day.After(today.Add(30*24*time.Hour)) {
		return []TimeSlot{}, nil // No slots for past or far-future dates
	}

	// We'll assume weekends are not available.
	if day.Weekday() == time.Saturday || day.Weekday() == time.Sunday {
		return []TimeSlot{}, nil
	}

	bookedSlots, err := s.store.GetBookingsForDay(day)
	if err != nil {
		return nil, err
	}

	var availableSlots []TimeSlot
	year, month, dayOfMonth := day.Date()
	loc := day.Location()

	// Generate all potential slots for the day
	slot := time.Date(year, month, dayOfMonth, s.cfg.WorkDayStartHour, 0, 0, 0, loc)
	endOfDay := time.Date(year, month, dayOfMonth, s.cfg.WorkDayEndHour, 0, 0, 0, loc)

	for slot.Before(endOfDay) {
		isBooked := false
		slotEnd := slot.Add(s.cfg.ConsultationDuration)

		for _, booked := range bookedSlots {
			if slot.Equal(booked.StartTime) {
				isBooked = true
				break
			}
		}

		if !isBooked && slot.After(time.Now()) {
			availableSlots = append(availableSlots, TimeSlot{
				StartTime: slot,
				EndTime:   slotEnd,
			})
		}
		slot = slotEnd
	}

	return availableSlots, nil
}

// CreateBooking creates a new booking and sends notifications.
func (s *BookingService) CreateBooking(req BookingRequest) (*Booking, error) {
	// In a real app, you'd add more validation here.
	// Is the start time on a 30-minute boundary? Is it within business hours?

	newBooking := Booking{
		ID:        uuid.NewString(),
		StartTime: req.StartTime,
		EndTime:   req.StartTime.Add(s.cfg.ConsultationDuration),
		Name:      req.Name,
		Email:     req.Email,
		Notes:     req.Notes,
	}

	if err := s.store.CreateBooking(newBooking); err != nil {
		return nil, err
	}

	// Send notifications concurrently to not block the API response.
	go func() {
		_ = s.emailer.SendBookingConfirmation(newBooking.Name, newBooking.Email, newBooking.StartTime)
		_ = s.emailer.SendBookingNotificationToAdmin(newBooking.Name, newBooking.Email, newBooking.StartTime, newBooking.Notes)
	}()

	return &newBooking, nil
}
