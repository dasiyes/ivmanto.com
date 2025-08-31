package config

import (
	"errors"
	"os"
	"time"
)

// Config holds all configuration for the application.
// It's a combination of the service-wide settings and specific feature configs.
type Config struct {
	Service ServiceMeta
	Booking BookingConfig
	Email   EmailConfig
}

// ServiceMeta holds general service information.
type ServiceMeta struct {
	Name      string `yaml:"name"`
	ProjectID string `yaml:"project_id"`
	Port      string `yaml:"port"`
}

// BookingConfig holds configuration specific to the booking service.
type BookingConfig struct {
	ConsultationDuration time.Duration
	WorkDayStartHour     int
	WorkDayEndHour       int
	CalendarID           string // Loaded from env var
}

// EmailConfig holds configuration for the email service.
type EmailConfig struct {
	SmtpHost      string `yaml:"smtp_host"`
	SmtpPort      string `yaml:"smtp_port"`
	SendFrom      string `yaml:"send_from"`
	SendFromAlias string `yaml:"send_from_alias"`
	SmtpPass      string // This MUST be loaded from a secret manager or env var
}

// Load returns the application configuration.
// In a real application, this would load from a file (e.g., config.yaml)
// or environment variables. For now, we'll hardcode the values.
func Load() (*Config, error) {
	cfg := &Config{
		Service: ServiceMeta{
			Name:      "ivmanto-backend-service",
			ProjectID: "ivmanto-com",
			Port:      os.Getenv("PORT"), // Use PORT env var from Cloud Run
		},
		Booking: BookingConfig{
			ConsultationDuration: 30 * time.Minute,
			WorkDayStartHour:     9,  // 9 AM UTC
			WorkDayEndHour:       17, // 5 PM UTC
			// This MUST be loaded from an environment variable.
			CalendarID: os.Getenv("CALENDAR_ID"),
		},
		Email: EmailConfig{
			SmtpHost:      "smtp.gmail.com",
			SmtpPort:      "587",
			SendFrom:      "nikolay.tonev@ivmanto.com",
			SendFromAlias: "accounts@ivmanto.com",
			// In a real app, load this securely.
			SmtpPass: os.Getenv("SMTP_PASS"),
		},
	}
	if cfg.Service.Port == "" {
		cfg.Service.Port = "8080" // Default for local dev
	}
	if cfg.Booking.CalendarID == "" {
		// We can't run without this.
		return nil, errors.New("CALENDAR_ID environment variable not set")
	}
	return cfg, nil
}
