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
	location   *time.Location
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

	// Fetch calendar details to get its timezone. This is crucial for correctly
	// interpreting date-only queries from the frontend.
	cal, err := srv.Calendars.Get(calendarID).Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve calendar details for ID %s: %w", calendarID, err)
	}
	loc, err := time.LoadLocation(cal.TimeZone)
	if err != nil {
		// Fallback to UTC if the location is not found, but log a warning.
		// This can happen in minimal container environments.
		fmt.Printf("WARNING: could not load timezone '%s', falling back to UTC. Error: %v\n", cal.TimeZone, err)
		loc = time.UTC
	}

	return &Service{
		calSvc:     srv,
		calendarID: calendarID,
		location:   loc,
	}, nil
}

// GetAvailability fetches available time slots for a given day.
func (s *Service) GetAvailability(day time.Time) ([]*calendar.Event, error) {
	// Use the calendar's timezone to define the start and end of the day.
	loc := s.location
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
	eventToUpdate.Summary = fmt.Sprintf("Consultation: %s", details.Name)
	// Add all relevant details to the description for the consultant to see.
	eventToUpdate.Description = fmt.Sprintf(
		"Client Name: %s\nClient Email: %s\n\nNotes:\n%s",
		details.Name,
		details.Email,
		details.Notes,
	)
	// The service account cannot invite external attendees without domain-wide delegation.
	// Instead of inviting them, we will send them a separate confirmation email with an .ics attachment.
	// This avoids the 403 Forbidden error.
	eventToUpdate.Attendees = nil

	// Add Google Meet conference
	// A unique RequestId prevents retries from creating duplicate meetings.
	conferenceRequestID := fmt.Sprintf("ivmanto-booking-%d", details.StartTime.Unix())
	eventToUpdate.ConferenceData = &calendar.ConferenceData{
		CreateRequest: &calendar.CreateConferenceRequest{
			RequestId: conferenceRequestID,
			// By leaving ConferenceSolutionKey empty, we let Google Calendar use the
			// default conference provider for the calendar, which is typically Google Meet.
			// This is more robust than explicitly requesting "hangoutsMeet".
		},
	}

	// Perform the update.
	// conferenceDataVersion: 1 tells Google to add a new Meet link.
	// SendUpdates: "none" because we are not sending calendar invites from here.
	updatedEvent, err := s.calSvc.Events.Update(s.calendarID, eventToUpdate.Id, eventToUpdate).
		ConferenceDataVersion(1).
		SendUpdates("none").
		Do()
	if err != nil {
		return nil, fmt.Errorf("unable to update event: %w", err)
	}

	return updatedEvent, nil
}

// Location returns the timezone of the calendar.
func (s *Service) Location() *time.Location {
	return s.location
}
