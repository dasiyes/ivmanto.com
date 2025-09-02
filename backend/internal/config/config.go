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
	AvailableSlotSummary string `env:"GCAL_AVAILABLE_SLOT_SUMMARY,required"` // <-- Add this line
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
			CalendarID:           os.Getenv("CALENDAR_ID"),
			AvailableSlotSummary: os.Getenv("GCAL_AVAILABLE_SLOT_SUMMARY"),
		},
		Email: EmailConfig{
			SmtpHost:      os.Getenv("SMTP_HOST"),
			SmtpPort:      os.Getenv("SMTP_PORT"),
			SendFrom:      os.Getenv("SEND_FROM"),
			SendFromAlias: os.Getenv("SEND_FROM_ALIAS"),
			// This is loaded from Secret Manager and passed as an env var.
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
	if cfg.Booking.AvailableSlotSummary == "" {
		return nil, errors.New("GCAL_AVAILABLE_SLOT_SUMMARY environment variable not set")
	}
	if cfg.Email.SmtpHost == "" {
		return nil, errors.New("SMTP_HOST environment variable not set")
	}
	if cfg.Email.SmtpPort == "" {
		return nil, errors.New("SMTP_PORT environment variable not set")
	}
	if cfg.Email.SendFrom == "" {
		return nil, errors.New("SEND_FROM environment variable not set")
	}
	if cfg.Email.SmtpPass == "" {
		return nil, errors.New("SMTP_PASS environment variable not set")
	}
	return cfg, nil
}
