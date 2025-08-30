package email

import (
	"log"
	"time"
)

// ContactMessage represents the data from the contact form.
type ContactMessage struct {
	Name    string
	Email   string
	Message string
}

// Service defines the interface for sending all application emails.
// This is the contract that the rest of the application will use,
// ensuring that high-level packages like 'booking' do not depend
// on the low-level implementation details of SMTP.
type Service interface {
	SendBookingConfirmation(name, toEmail string, startTime time.Time) error
	SendBookingNotificationToAdmin(name, clientEmail string, startTime time.Time, notes string) error
	SendContactMessage(msg ContactMessage) error
}

// MockService is a no-op implementation of the Service interface for development.
type MockService struct{}

// NewMockService creates a new mock email service.
func NewMockService() *MockService {
	return &MockService{}
}

func (s *MockService) SendBookingConfirmation(name, toEmail string, startTime time.Time) error {
	log.Printf("SIMULATING: Sending booking confirmation to %s for %s", toEmail, startTime)
	return nil
}

func (s *MockService) SendBookingNotificationToAdmin(name, clientEmail string, startTime time.Time, notes string) error {
	log.Printf("SIMULATING: Sending admin notification for booking with %s (%s) at %s", name, clientEmail, startTime)
	return nil
}

func (s *MockService) SendContactMessage(msg ContactMessage) error {
	log.Printf("SIMULATING: Sending contact message from %s <%s>: %s", msg.Name, msg.Email, msg.Message)
	return nil
}
