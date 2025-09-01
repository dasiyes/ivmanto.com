package gcal

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

var (
	// ErrSlotNotFound is returned when a requested booking slot cannot be found or is already booked.
	ErrSlotNotFound = errors.New("slot not found or already booked")
)

// Service interacts with the Google Calendar API.
type Service struct {
	calSvc               *calendar.Service
	calendarID           string
	location             *time.Location
	availableSlotSummary string
}

// BookingDetails contains information for a new booking.
type BookingDetails struct {
	StartTime time.Time
	Name      string
	Email     string
	Notes     string
}

// NewService creates a new calendar service client.
func NewService(ctx context.Context, credentialFile, calendarID, availableSlotSummary string) (*Service, error) {
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
		calSvc:               srv,
		calendarID:           calendarID,
		location:             loc,
		availableSlotSummary: availableSlotSummary,
	}, nil
}

// GetAvailability fetches available time slots for a given day.
// It now lists events with a specific summary to find available slots.
func (s *Service) GetAvailability(day time.Time) ([]*calendar.Event, error) {
	loc := s.location
	startOfDay := time.Date(day.Year(), day.Month(), day.Day(), 0, 0, 0, 0, loc)
	endOfDay := startOfDay.Add(24 * time.Hour)

	events, err := s.calSvc.Events.List(s.calendarID).
		TimeMin(startOfDay.Format(time.RFC3339)).
		TimeMax(endOfDay.Format(time.RFC3339)).
		Q(s.availableSlotSummary). // Search for events with the "Available" summary
		SingleEvents(true).
		OrderBy("startTime").
		Do()
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve availability events: %w", err)
	}

	// The original implementation returned []*calendar.TimePeriod, but the handler
	// expects []*calendar.Event. We will return the events directly.
	return events.Items, nil
}

// BookSlot books a consultation by finding an "Available" event and updating it.
// This provides an atomic way to claim a slot.
func (s *Service) BookSlot(details BookingDetails) (*calendar.Event, error) {
	// 1. Find the specific "Available" event at the requested start time.
	// We query a very narrow time window to ensure we get the exact slot.
	// Using Q (text search) is a good filter, but we'll double-check the summary.
	list, err := s.calSvc.Events.List(s.calendarID).
		TimeMin(details.StartTime.Format(time.RFC3339)).
		TimeMax(details.StartTime.Add(1 * time.Second).Format(time.RFC3339)).
		Q(s.availableSlotSummary).
		SingleEvents(true).
		MaxResults(1).
		Do()

	if err != nil {
		return nil, fmt.Errorf("unable to retrieve event to book: %w", err)
	}

	if len(list.Items) == 0 {
		return nil, ErrSlotNotFound
	}

	eventToBook := list.Items[0]

	// 2. Verify the event is indeed an available slot and not already booked.
	if eventToBook.Summary != s.availableSlotSummary {
		return nil, ErrSlotNotFound
	}

	// Generate a unique cancellation token for this booking.
	cancellationUUID, err := uuid.NewRandom()
	if err != nil {
		// This is a server-side issue, but we shouldn't fail the whole booking for it.
		// Log it and continue. The user just won't get a cancellation link.
		log.Printf("WARNING: could not generate cancellation token UUID: %v", err)
	} else {
		cancellationToken := cancellationUUID.String()
		if eventToBook.ExtendedProperties == nil {
			eventToBook.ExtendedProperties = &calendar.EventExtendedProperties{}
		}
		if eventToBook.ExtendedProperties.Private == nil {
			eventToBook.ExtendedProperties.Private = make(map[string]string)
		}
		// Store booking details for later use (e.g., cancellation notifications).
		eventToBook.ExtendedProperties.Private["cancellation_token"] = cancellationToken
		eventToBook.ExtendedProperties.Private["client_name"] = details.Name
		eventToBook.ExtendedProperties.Private["client_email"] = details.Email
	}

	// 3. Update the event with the client's details.
	eventToBook.Summary = fmt.Sprintf("Consultation: %s", details.Name)
	eventToBook.Description = fmt.Sprintf(
		"Client Name: %s\nClient Email: %s\n\nNotes:\n%s",
		details.Name,
		details.Email,
		details.Notes,
	)
	// By default, a service account cannot add attendees to an event without
	// being granted domain-wide delegation of authority. We will handle sending
	// a confirmation email from our own backend instead.
	eventToBook.Attendees = nil
	// Request Google Meet conference data to be added to the event.
	eventToBook.ConferenceData = &calendar.ConferenceData{
		CreateRequest: &calendar.CreateConferenceRequest{
			RequestId: fmt.Sprintf("ivmanto-booking-%d", time.Now().UnixNano()),
			// By omitting the ConferenceSolutionKey, we ask Google Calendar to use
			// the default conference provider configured for this calendar. This is
			// the most robust method for Workspace accounts, as it respects the
			// server-side configuration instead of being overly prescriptive.
			// This resolves the "Invalid conference type value" error.
		},
	}

	// 4. Atomically update the event. The ETag mechanism handled by the client library
	// ensures that if the event was changed between our read and write, this will fail.
	_, err = s.calSvc.Events.Update(s.calendarID, eventToBook.Id, eventToBook).
		ConferenceDataVersion(1). // Required when modifying conference data
		// We remove SendUpdates("all") as there are no attendees to notify.
		Do()

	if err != nil {
		// Check for a 409 Conflict or 412 Precondition Failed, which indicates the slot was just taken.
		if gerr, ok := err.(*googleapi.Error); ok && (gerr.Code == http.StatusConflict || gerr.Code == http.StatusPreconditionFailed) {
			return nil, ErrSlotNotFound
		}
		return nil, fmt.Errorf("failed to update event during booking: %w", err)
	}

	// 5. Re-fetch the event to ensure we have the latest data, including the generated HangoutLink.
	// The conference generation can be asynchronous, and the object returned from Update might not have it.
	finalEvent, err := s.calSvc.Events.Get(s.calendarID, eventToBook.Id).Do()
	if err != nil {
		// If we can't get the event after updating it, it's a problem, but the booking was made.
		// We'll log it and return an error, but a more robust system might handle this differently.
		return nil, fmt.Errorf("event was booked but failed to retrieve final details: %w", err)
	}

	return finalEvent, nil
}

// CancelSlot finds an event by its cancellation token and reverts it to an "Available" slot.
func (s *Service) CancelSlot(token string) (*calendar.Event, error) {
	if token == "" {
		return nil, errors.New("cancellation token cannot be empty")
	}

	// 1. Find the event using the private extended property.
	// This is the most secure way to identify the event to cancel.
	list, err := s.calSvc.Events.List(s.calendarID).
		PrivateExtendedProperty(fmt.Sprintf("cancellation_token=%s", token)).
		MaxResults(1).
		Do()

	if err != nil {
		return nil, fmt.Errorf("failed to search for event by cancellation token: %w", err)
	}

	if len(list.Items) == 0 {
		return nil, ErrSlotNotFound // Token is invalid or event already cancelled/deleted.
	}

	eventToCancel := list.Items[0]

	// 2. Revert the event to an "Available" slot.
	eventToCancel.Summary = s.availableSlotSummary
	eventToCancel.Description = "" // Clear client notes
	eventToCancel.Attendees = nil
	// To remove conference data, we must set it to an empty object and
	// specify ConferenceDataVersion(1) in the update call. Setting it to nil
	// is not sufficient.
	eventToCancel.ConferenceData = &calendar.ConferenceData{}
	eventToCancel.ExtendedProperties.Private["cancellation_token"] = "" // Invalidate the token

	// 3. Update the event.
	cancelledEvent, err := s.calSvc.Events.Update(s.calendarID, eventToCancel.Id, eventToCancel).
		ConferenceDataVersion(1). // Required when modifying conference data.
		Do()
	if err != nil {
		return nil, fmt.Errorf("failed to update event during cancellation: %w", err)
	}

	return cancelledEvent, nil
}

// Location returns the timezone of the calendar.
func (s *Service) Location() *time.Location {
	return s.location
}
