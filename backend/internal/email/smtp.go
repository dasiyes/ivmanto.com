package email

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
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

// stripPlusAlias removes the +alias part of an email address, which is sometimes
// required for the SMTP RCPT TO command.
// e.g., "user+alias@example.com" becomes "user@example.com".
func stripPlusAlias(email string) string {
	parts := strings.SplitN(email, "@", 2)
	if len(parts) != 2 {
		return email // Not a standard email format, return as is.
	}
	localPart, domainPart := parts[0], parts[1]

	plusIndex := strings.Index(localPart, "+")
	if plusIndex == -1 {
		return email // No plus alias, return as is.
	}

	return fmt.Sprintf("%s@%s", localPart[:plusIndex], domainPart)
}

// send is a private helper to construct and dispatch an email.
func (s *SmtpService) send(to, cc []string, subject, body string) error {
	// Construct the email message headers.
	headers := make(map[string]string)
	headers["From"] = fmt.Sprintf("%s <%s>", s.cfg.SendFromAlias, s.cfg.SendFrom)
	headers["To"] = strings.Join(to, ", ")
	if len(cc) > 0 {
		headers["Cc"] = strings.Join(cc, ", ")
	}
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

	allRecipients := append(to, cc...)

	log.Printf("Connecting to SMTP server at %s", addr)
	c, err := smtp.Dial(addr)
	if err != nil {
		log.Printf("ERROR: Failed to connect to SMTP server: %v", err)
		return err
	}
	defer c.Close()

	// It's good practice to say hello to the server.
	// The hostname can be anything, but localhost is common for clients.
	if err = c.Hello("localhost"); err != nil {
		log.Printf("ERROR: Failed to send HELO to SMTP server: %v", err)
		return err
	}

	// Use STARTTLS. Gmail requires this on port 587.
	if ok, _ := c.Extension("STARTTLS"); ok {
		log.Println("Server supports STARTTLS. Upgrading connection...")
		config := &tls.Config{ServerName: s.cfg.SmtpHost}
		if err = c.StartTLS(config); err != nil {
			log.Printf("ERROR: Failed to start TLS: %v", err)
			return err
		}
	}

	// Authenticate.
	if s.auth != nil {
		log.Println("Authenticating...")
		if err = c.Auth(s.auth); err != nil {
			log.Printf("ERROR: Authentication failed: %v", err)
			return err
		}
	}

	// Set the sender.
	log.Printf("Setting sender to %s", s.cfg.SendFrom)
	if err = c.Mail(s.cfg.SendFrom); err != nil {
		log.Printf("ERROR: Failed to set sender: %v", err)
		return err
	}

	// Set the recipients.
	for _, rcpt := range allRecipients {
		// Strip +alias for the RCPT TO command, as some servers require the base address.
		baseRcpt := stripPlusAlias(rcpt)
		log.Printf("Adding recipient %s (sending to %s)", rcpt, baseRcpt)
		if err = c.Rcpt(baseRcpt); err != nil {
			log.Printf("ERROR: Failed to add recipient %s (as %s): %v", rcpt, baseRcpt, err)
			return err
		}
	}

	// Get the writer for the data and write the message.
	log.Println("Sending email body...")
	wc, err := c.Data()
	if err != nil {
		log.Printf("ERROR: Failed to get data writer: %v", err)
		return err
	}
	_, err = wc.Write(msg.Bytes())
	if err != nil {
		log.Printf("ERROR: Failed to write email body: %v", err)
		return err
	}
	err = wc.Close()
	if err != nil {
		log.Printf("ERROR: Failed to close data writer: %v", err)
		return err
	}

	// Quit the session.
	log.Println("Quitting SMTP session.")
	err = c.Quit()
	if err != nil {
		log.Printf("ERROR: Failed to quit SMTP session: %v", err)
		return err
	}

	log.Printf("Email sent successfully to %s", strings.Join(allRecipients, ", "))
	return nil
}

// SendBookingConfirmation sends a confirmation email to the user.
func (s *SmtpService) SendBookingConfirmation(name, toEmail string, startTime time.Time) error {
	subject := "Your consultation is confirmed!"
	// In a real app, this body would come from an HTML template.
	body := fmt.Sprintf("Hi %s,<br><br>Your consultation on %s is confirmed.<br><br>Thanks,<br>The Team", name, startTime.Format(time.RFC1123))
	return s.send([]string{toEmail}, nil, subject, body)
}

// SendBookingNotificationToAdmin sends a notification email to the admin.
func (s *SmtpService) SendBookingNotificationToAdmin(name, clientEmail string, startTime time.Time, notes string) error {
	// Use a '+booking' alias to ensure delivery to the admin's inbox.
	parts := strings.Split(s.cfg.SendFrom, "@")
	var adminEmail string
	if len(parts) == 2 {
		adminEmail = fmt.Sprintf("%s+booking@%s", parts[0], parts[1])
	} else {
		adminEmail = s.cfg.SendFrom // Fallback for non-standard emails
	}

	subject := "New Consultation Booked!"
	body := fmt.Sprintf("New booking with:<br>Name: %s<br>Email: %s<br>Time: %s<br>Notes: %s", name, clientEmail, startTime.Format(time.RFC1123), notes)
	return s.send([]string{adminEmail}, nil, subject, body)
}

// SendContactMessage sends the contact form message to the admin.
func (s *SmtpService) SendContactMessage(msg ContactMessage) error {
	// Trick to help Gmail deliver to inbox: add a suffix to the username.
	// e.g., nikolay.tonev@ivmanto.com -> nikolay.tonev+contact@ivmanto.com
	parts := strings.Split(s.cfg.SendFrom, "@")
	var adminEmail string
	if len(parts) == 2 {
		adminEmail = fmt.Sprintf("%s+contact@%s", parts[0], parts[1])
	} else {
		adminEmail = s.cfg.SendFrom // Fallback for non-standard emails
	}

	subject := fmt.Sprintf("New Contact Message from %s", msg.Name)
	body := fmt.Sprintf("From: %s &lt;%s&gt;<br><br>Message:<br>%s", msg.Name, msg.Email, msg.Message)

	var ccList []string
	if msg.SendCopyToSelf {
		ccList = append(ccList, msg.Email)
	}

	return s.send([]string{adminEmail}, ccList, subject, body)
}
