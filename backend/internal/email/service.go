package email

import "time"

// Service defines the interface for sending all types of application emails.
type Service interface {
	SendContactMessage(msg ContactMessage) error
	SendBookingConfirmation(details BookingConfirmationDetails) error
	SendBookingNotificationToAdmin(name, clientEmail string, startTime time.Time, notes string) error
	SendBookingCancellationToClient(toName, toEmail string, startTime time.Time) error
	SendBookingCancellationToAdmin(clientName, clientEmail string, startTime time.Time) error
	SendGeneratedIdeas(toEmail, topic string, ideasBody string) error
}
