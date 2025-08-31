package gcal

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

const (
	// You will create events in your Google Calendar with this summary
	availableSlotSummary = "Available for Consultation"
)

var (
	// ErrSlotNotFound is returned when a requested booking slot cannot be found or is already booked.
	ErrSlotNotFound = errors.New("slot not found or already booked")
)

// Service interacts with the Google Calendar API.
type Service struct {
	calSvc     *calendar.Service
	calendarID string
}

// BookingDetails contains information for a new booking.
type BookingDetails struct {
	StartTime time.Time
	Name      string
	Email     string
	Notes     string
}

// NewService creates a new calendar service client.
func NewService(ctx context.Context, credentialFile, calendarID string) (*Service, error) {
	b, err := os.ReadFile(credentialFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read client secret file '%s': %w", credentialFile, err)
	}

	// We use a service account for server-to-server authentication.
	config, err := google.JWTConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse client secret file to config: %w", err)
	}
	client := config.Client(ctx)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Calendar client: %w", err)
	}

	return &Service{
		calSvc:     srv,
		calendarID: calendarID,
	}, nil
}

// GetAvailability fetches available time slots for a given day.
func (s *Service) GetAvailability(day time.Time) ([]*calendar.Event, error) {
	// We need to specify the timezone of the calendar to get correct results.
	// For now, let's assume UTC. In a future version, this could be configurable.
	loc := time.UTC
	startOfDay := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, loc)
	endOfDay := startOfDay.Add(24 * time.Hour)

	events, err := s.calSvc.Events.List(s.calendarID).
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(startOfDay.Format(time.RFC3339)).
		TimeMax(endOfDay.Format(time.RFC3339)).
		Q(availableSlotSummary). // Filter for events with our specific summary
		OrderBy("startTime").Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve events from calendar: %w", err)
	}

	return events.Items, nil
}

// BookSlot books a consultation and adds a Google Meet link.
func (s *Service) BookSlot(details BookingDetails) (*calendar.Event, error) {
	// Find the specific "Available" event to convert into a booking.
	// We search a 1-minute window around the start time to be precise and atomic.
	timeMin := details.StartTime.Format(time.RFC3339)
	timeMax := details.StartTime.Add(1 * time.Minute).Format(time.RFC3339)

	events, err := s.calSvc.Events.List(s.calendarID).
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(timeMin).
		TimeMax(timeMax).
		Q(availableSlotSummary).
		MaxResults(1).
		Do()
	if err != nil {
		return nil, fmt.Errorf("unable to find available slot: %w", err)
	}

	if len(events.Items) == 0 {
		return nil, ErrSlotNotFound
	}

	eventToUpdate := events.Items[0]

	// Update event details
	eventToUpdate.Summary = fmt.Sprintf("Booked: Consultation with %s", details.Name)
	eventToUpdate.Description = details.Notes
	eventToUpdate.Attendees = []*calendar.EventAttendee{
		{Email: details.Email},
	}

	// Add Google Meet conference
	// A unique RequestId prevents retries from creating duplicate meetings.
	conferenceRequestID := fmt.Sprintf("ivmanto-booking-%d", details.StartTime.Unix())
	eventToUpdate.ConferenceData = &calendar.ConferenceData{
		CreateRequest: &calendar.CreateConferenceRequest{
			RequestId: conferenceRequestID,
			ConferenceSolutionKey: &calendar.ConferenceSolutionKey{
				Type: "hangoutsMeet",
			},
		},
	}

	// Perform the update.
	// conferenceDataVersion: 1 tells Google to add a new Meet link.
	// SendUpdates: "all" sends calendar invites to new attendees.
	updatedEvent, err := s.calSvc.Events.Update(s.calendarID, eventToUpdate.Id, eventToUpdate).
		ConferenceDataVersion(1).
		SendUpdates("all").
		Do()
	if err != nil {
		return nil, fmt.Errorf("unable to update event: %w", err)
	}

	return updatedEvent, nil
}
