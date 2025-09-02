package gcal

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
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

// Service defines the interface for interacting with Google Calendar.
type Service interface {
	GetAvailability(day time.Time) ([]*calendar.Event, error)
	BookSlot(details BookingDetails) (*calendar.Event, error)
	CancelBooking(ctx context.Context, token string) (*calendar.Event, error)
	Location() *time.Location
}

// gcalService implements the Service interface for Google Calendar.
type gcalService struct {
	calSvc               *calendar.Service
	calendarID           string
	location             *time.Location
	availableSlotSummary string
}

// BookingDetails contains information for a new booking.
type BookingDetails struct {
	EventID string
	Name    string
	Email   string
	Notes   string
}

// NewService creates a new calendar service client.
func NewService(ctx context.Context, calendarID, availableSlotSummary string) (Service, error) {
	// When running on Google Cloud (like Cloud Run), the client library will
	// automatically find the credentials of the service account the service is
	// running as. This is the recommended and most secure way to authenticate.
	// For local development, it uses the credentials from `gcloud auth application-default login`.
	creds, err := google.FindDefaultCredentials(ctx, calendar.CalendarScope)
	if err != nil {
		return nil, fmt.Errorf("unable to find default credentials: %w", err)
	}

	srv, err := calendar.NewService(ctx, option.WithCredentials(creds))
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

	return &gcalService{
		calSvc:               srv,
		calendarID:           calendarID,
		location:             loc,
		availableSlotSummary: strings.TrimSpace(availableSlotSummary),
	}, nil
}

// GetAvailability fetches available time slots for a given day.
// It now lists events with a specific summary to find available slots.
func (s *gcalService) GetAvailability(day time.Time) ([]*calendar.Event, error) {
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
func (s *gcalService) BookSlot(details BookingDetails) (*calendar.Event, error) {
	log.Printf("INFO: [gcal] Attempting to book event with ID: %s", details.EventID)

	// 1. Get the event directly by its unique ID. This is more reliable than searching.
	eventToBook, err := s.calSvc.Events.Get(s.calendarID, details.EventID).Do()
	if err != nil {
		// If the error is 404, it means the event doesn't exist, which we treat as a slot not found.
		if gerr, ok := err.(*googleapi.Error); ok && gerr.Code == http.StatusNotFound {
			return nil, ErrSlotNotFound
		}
		return nil, fmt.Errorf("unable to retrieve event to book with ID %s: %w", details.EventID, err)
	}

	log.Printf("INFO: [gcal] Found available event %s to book.", eventToBook.Id)

	// 2. Verify the event is indeed an available slot and not already booked.
	// We trim the space from the calendar summary to be robust against accidental whitespace.
	if strings.TrimSpace(eventToBook.Summary) != s.availableSlotSummary {
		log.Printf("ERROR: [gcal] Slot verification failed. Event summary from calendar: '%s' does not match expected summary from config: '%s'", strings.TrimSpace(eventToBook.Summary), s.availableSlotSummary)
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
	// We do not add attendees directly, as this requires domain-wide delegation for the service account.
	// Instead, our application will send a confirmation email with an .ics attachment.
	eventToBook.Attendees = nil
	// Request Google Meet conference data to be added to the event.
	eventToBook.ConferenceData = &calendar.ConferenceData{
		CreateRequest: &calendar.CreateConferenceRequest{
			RequestId: fmt.Sprintf("ivmanto-booking-%d", time.Now().UnixNano()),
			// For a Google Workspace account, "hangoutsMeet" is the modern and
			// recommended conference solution type. This ensures a unique
			// Google Meet link is generated for every booking, which is exactly
			// what we want.
			ConferenceSolutionKey: &calendar.ConferenceSolutionKey{Type: "hangoutsMeet"},
		},
	}

	// 4. Atomically update the event. The ETag mechanism handled by the client library
	// ensures that if the event was changed between our read and write, this will fail.
	updatedEvent, err := s.calSvc.Events.Update(s.calendarID, eventToBook.Id, eventToBook).
		ConferenceDataVersion(1). // Required when modifying conference data.
		Do()

	if err != nil {
		// Check for a 409 Conflict or 412 Precondition Failed, which indicates the slot was just taken.
		if gerr, ok := err.(*googleapi.Error); ok && (gerr.Code == http.StatusConflict || gerr.Code == http.StatusPreconditionFailed) {
			return nil, ErrSlotNotFound
		}
		return nil, fmt.Errorf("failed to update event during booking: %w", err)
	}
	log.Printf("INFO: [gcal] Successfully updated event %s with booking details.", updatedEvent.Id)

	// 5. Re-fetch the event to ensure we have the latest data.
	// Conference data generation can be asynchronous. We'll retry a few times to get it.
	var finalEvent *calendar.Event
	var getErr error
	maxRetries := 3
	retryDelay := 500 * time.Millisecond

	for i := 0; i < maxRetries; i++ {
		log.Printf("INFO: [gcal] Fetching event %s to get conference data (Attempt %d/%d)...", updatedEvent.Id, i+1, maxRetries)
		finalEvent, getErr = s.calSvc.Events.Get(s.calendarID, updatedEvent.Id).Do()
		if getErr != nil {
			log.Printf("ERROR: [gcal] Failed to fetch event %s: %v", updatedEvent.Id, getErr)
			return nil, getErr
		}

		// Check if conference data is available. We check both HangoutLink and ConferenceData.EntryPoints.
		if finalEvent.HangoutLink != "" || (finalEvent.ConferenceData != nil && len(finalEvent.ConferenceData.EntryPoints) > 0) {
			log.Printf("INFO: [gcal] Successfully fetched event with conference data.")
			return finalEvent, nil
		}

		if i < maxRetries-1 {
			log.Printf("WARN: [gcal] Conference data not yet available for event %s. Retrying in %v...", updatedEvent.Id, retryDelay)
			time.Sleep(retryDelay)
		}
	}

	// If we exit the loop, it means we couldn't get the link after several retries.
	// We'll return the last fetched event, and the email service will have to handle an empty link.
	log.Printf("WARN: [gcal] Could not retrieve conference data for event %s after %d retries. Proceeding without it.", updatedEvent.Id, maxRetries)
	return finalEvent, nil
}

// CancelBooking finds an event by its cancellation token and reverts it to an available slot.
// It returns the original event details for notification purposes.
func (s *gcalService) CancelBooking(ctx context.Context, token string) (*calendar.Event, error) {
	log.Printf("INFO: [gcal] Searching for event with cancellation token %s...", token[:8])
	// 1. Find the event using the private extended property.
	query := fmt.Sprintf("cancellation_token=%s", token)
	events, err := s.calSvc.Events.List(s.calendarID).
		PrivateExtendedProperty(query).
		MaxResults(1).
		Do()
	if err != nil {
		return nil, fmt.Errorf("failed to query for event with token: %w", err)
	}

	if len(events.Items) == 0 {
		log.Printf("WARN: [gcal] No event found for cancellation token %s...", token[:8])
		return nil, ErrSlotNotFound // Using existing error for "not found"
	}

	eventToCancel := events.Items[0]
	log.Printf("INFO: [gcal] Found event %s to cancel.", eventToCancel.Id)

	// 2. Preserve original details for notifications before modifying.
	// We retrieve the client details from the private properties we stored during booking.
	originalEvent := &calendar.Event{
		Id:      eventToCancel.Id,
		Summary: eventToCancel.Summary,
		Start:   eventToCancel.Start,
		End:     eventToCancel.End,
		Attendees: []*calendar.EventAttendee{
			{
				DisplayName: eventToCancel.ExtendedProperties.Private["client_name"],
				Email:       eventToCancel.ExtendedProperties.Private["client_email"],
			},
		},
	}

	// 3. Update the event to revert it to an "Available" slot.
	eventToCancel.Summary = s.availableSlotSummary
	eventToCancel.Description = "This slot is now available for booking."
	eventToCancel.Attendees = nil
	// Set ConferenceData to nil to explicitly remove the Google Meet link.
	// This requires ConferenceDataVersion(1) in the Update call.
	eventToCancel.ConferenceData = nil
	delete(eventToCancel.ExtendedProperties.Private, "cancellation_token")
	delete(eventToCancel.ExtendedProperties.Private, "client_name")
	delete(eventToCancel.ExtendedProperties.Private, "client_email")

	// 4. Persist the update to Google Calendar.
	_, err = s.calSvc.Events.Update(s.calendarID, eventToCancel.Id, eventToCancel).ConferenceDataVersion(1).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to update event to available: %w", err)
	}
	log.Printf("INFO: [gcal] Successfully reverted event %s to an available slot.", eventToCancel.Id)

	return originalEvent, nil
}

// Location returns the timezone of the calendar.
func (s *gcalService) Location() *time.Location {
	return s.location
}
