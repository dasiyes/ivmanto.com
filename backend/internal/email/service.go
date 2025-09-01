package email

import (
	"time"

	"google.golang.org/api/calendar/v3"
)

// ContactMessage represents the data from the contact form.
type ContactMessage struct {
	Name           string
	Email          string
	Message        string
	SendCopyToSelf bool
}

// Service defines the interface for sending emails.
type Service interface {
	SendContactMessage(msg ContactMessage) error
	SendBookingConfirmation(toName, toEmail string, event *calendar.Event) error
	SendBookingNotificationToAdmin(name, clientEmail string, startTime time.Time, notes string) error
	SendBookingCancellationToClient(toName, toEmail string, startTime time.Time) error
	SendBookingCancellationToAdmin(clientName, clientEmail string, startTime time.Time) error
}
