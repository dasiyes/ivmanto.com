package config

import (
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/compute/metadata"
)

// Config holds all configuration for the application.
type Config struct {
	Service   ServiceConfig
	Email     EmailConfig
	GCal      GCalConfig
	GCP       GCPConfig
	Ideas     IdeasConfig
	Analytics AnalyticsConfig
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

// GCPConfig holds project-level Google Cloud configuration.
type GCPConfig struct {
	ProjectID string
	Location  string
}

// Add a new struct for Ideas configuration
type IdeasConfig struct {
	GenerateIdeasPromptTemplate string `env:"GENERATE_IDEAS_PROMPT_TEMPLATE"`
}

// Add this struct
type AnalyticsConfig struct {
	ApiSecret     string `env:"GA_API_SECRET,required"`
	MeasurementID string `env:"GA_MEASUREMENT_ID,required"`
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	var missingVars []string

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	smtpHost := os.Getenv("SMTP_HOST")
	if smtpHost == "" {
		missingVars = append(missingVars, "SMTP_HOST")
	}
	smtpPort := os.Getenv("SMTP_PORT")
	if smtpPort == "" {
		missingVars = append(missingVars, "SMTP_PORT")
	}
	sendFrom := os.Getenv("SEND_FROM")
	if sendFrom == "" {
		missingVars = append(missingVars, "SEND_FROM")
	}
	sendFromAlias := os.Getenv("SEND_FROM_ALIAS")
	smtpPass := os.Getenv("SMTP_PASS")
	if smtpPass == "" {
		missingVars = append(missingVars, "SMTP_PASS")
	}

	calendarID := os.Getenv("CALENDAR_ID")
	if calendarID == "" {
		missingVars = append(missingVars, "CALENDAR_ID")
	}
	availableSlotSummary := os.Getenv("GCAL_AVAILABLE_SLOT_SUMMARY")
	if availableSlotSummary == "" {
		missingVars = append(missingVars, "GCAL_AVAILABLE_SLOT_SUMMARY")
	}

	projectID, err := metadata.ProjectID()
	if err != nil {
		projectID = os.Getenv("GCP_PROJECT_ID")
	}
	if projectID == "" {
		missingVars = append(missingVars, "GCP_PROJECT_ID")
	}

	location := os.Getenv("GCP_LOCATION")
	if location == "" {
		missingVars = append(missingVars, "GCP_LOCATION")
	}

	generateIdeasPromptTemplate := os.Getenv("GENERATE_IDEAS_PROMPT_TEMPLATE")

	// Load Analytics config
	gaApiSecret := os.Getenv("GA_API_SECRET")
	if gaApiSecret == "" {
		missingVars = append(missingVars, "GA_API_SECRET")
	}
	gaMeasurementID := os.Getenv("GA_MEASUREMENT_ID")
	if gaMeasurementID == "" {
		missingVars = append(missingVars, "GA_MEASUREMENT_ID")
	}

	if len(missingVars) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %s", strings.Join(missingVars, ", "))
	}

	return &Config{
		Service: ServiceConfig{Port: port},
		Email:   EmailConfig{SmtpHost: smtpHost, SmtpPort: smtpPort, SendFrom: sendFrom, SendFromAlias: sendFromAlias, SmtpPass: smtpPass},
		GCal:    GCalConfig{CalendarID: calendarID, AvailableSlotSummary: availableSlotSummary},
		GCP:     GCPConfig{ProjectID: projectID, Location: location},
		Ideas:   IdeasConfig{GenerateIdeasPromptTemplate: generateIdeasPromptTemplate},
		Analytics: AnalyticsConfig{
			ApiSecret:     gaApiSecret,
			MeasurementID: gaMeasurementID,
		},
	}, nil
}
