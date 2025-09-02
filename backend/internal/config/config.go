package config

import (
	"fmt"
	"os"
)

// Config holds all configuration for the application.
type Config struct {
	Service ServiceConfig
	Email   EmailConfig
	GCal    GCalConfig
}

// ServiceConfig holds configuration for the HTTP service.
type ServiceConfig struct {
	Port string
}

// EmailConfig holds configuration for the SMTP email service.
type EmailConfig struct {
	SmtpHost      string
	SmtpPort      string
	SendFrom      string
	SendFromAlias string
	SmtpPass      string // Loaded from Secret Manager
}

// GCalConfig holds configuration for the Google Calendar service.
type GCalConfig struct {
	CalendarID           string
	AvailableSlotSummary string
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	sendFrom := os.Getenv("SEND_FROM")
	sendFromAlias := os.Getenv("SEND_FROM_ALIAS")
	smtpPass := os.Getenv("SMTP_PASS")

	if smtpHost == "" || smtpPort == "" || sendFrom == "" || smtpPass == "" {
		return nil, fmt.Errorf("one or more required email environment variables are not set (SMTP_HOST, SMTP_PORT, SEND_FROM, SMTP_PASS)")
	}

	calendarID := os.Getenv("CALENDAR_ID")
	availableSlotSummary := os.Getenv("GCAL_AVAILABLE_SLOT_SUMMARY")
	if calendarID == "" || availableSlotSummary == "" {
		return nil, fmt.Errorf("one or more required Google Calendar environment variables are not set (CALENDAR_ID, GCAL_AVAILABLE_SLOT_SUMMARY)")
	}

	return &Config{
		Service: ServiceConfig{Port: port},
		Email:   EmailConfig{SmtpHost: smtpHost, SmtpPort: smtpPort, SendFrom: sendFrom, SendFromAlias: sendFromAlias, SmtpPass: smtpPass},
		GCal:    GCalConfig{CalendarID: calendarID, AvailableSlotSummary: availableSlotSummary},
	}, nil
}
