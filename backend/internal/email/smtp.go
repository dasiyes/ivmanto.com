package email

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strings"
	"time"

	"ivmanto.com/backend/internal/config"
)

// SmtpService is a concrete implementation of the email Service using SMTP.
type SmtpService struct {
	cfg  *config.EmailConfig
	auth smtp.Auth
}

// NewSmtpService creates a new SMTP email service.
// It requires a valid EmailConfig. The SMTP password should be loaded from a secure source.
func NewSmtpService(cfg *config.EmailConfig) *SmtpService {
	// In a real app, you'd check if cfg.SmtpPass is empty and handle it.
	auth := smtp.PlainAuth("", cfg.SendFrom, cfg.SmtpPass, cfg.SmtpHost)
	return &SmtpService{
		cfg:  cfg,
		auth: auth,
	}
}

// send is a private helper to construct and dispatch an email.
func (s *SmtpService) send(to []string, subject, body string) error {
	// Construct the email message headers.
	headers := make(map[string]string)
	headers["From"] = fmt.Sprintf("%s <%s>", s.cfg.SendFromAlias, s.cfg.SendFrom)
	headers["To"] = strings.Join(to, ", ")
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"utf-8\""

	// Build the full message.
	var msg bytes.Buffer
	for k, v := range headers {
		msg.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	msg.WriteString("\r\n")
	msg.WriteString(body)

	addr := fmt.Sprintf("%s:%s", s.cfg.SmtpHost, s.cfg.SmtpPort)

	// Send the email.
	return smtp.SendMail(addr, s.auth, s.cfg.SendFrom, to, msg.Bytes())
}

// SendBookingConfirmation sends a confirmation email to the user.
func (s *SmtpService) SendBookingConfirmation(name, toEmail string, startTime time.Time) error {
	subject := "Your consultation is confirmed!"
	// In a real app, this body would come from an HTML template.
	body := fmt.Sprintf("Hi %s,<br><br>Your consultation on %s is confirmed.<br><br>Thanks,<br>The Team", name, startTime.Format(time.RFC1123))
	return s.send([]string{toEmail}, subject, body)
}

// SendBookingNotificationToAdmin sends a notification email to the admin.
func (s *SmtpService) SendBookingNotificationToAdmin(name, clientEmail string, startTime time.Time, notes string) error {
	// The admin's email should also be in the config.
	adminEmail := s.cfg.SendFrom // Sending to self for now.
	subject := "New Consultation Booked!"
	body := fmt.Sprintf("New booking with:<br>Name: %s<br>Email: %s<br>Time: %s<br>Notes: %s", name, clientEmail, startTime.Format(time.RFC1123), notes)
	return s.send([]string{adminEmail}, subject, body)
}

// SendContactMessage sends the contact form message to the admin.
func (s *SmtpService) SendContactMessage(msg ContactMessage) error {
	adminEmail := s.cfg.SendFrom // Sending to self.
	subject := fmt.Sprintf("New Contact Message from %s", msg.Name)
	body := fmt.Sprintf("From: %s &lt;%s&gt;<br><br>Message:<br>%s", msg.Name, msg.Email, msg.Message)
	return s.send([]string{adminEmail}, subject, body)
}
