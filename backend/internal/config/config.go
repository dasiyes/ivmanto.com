package config

import (
	"fmt"
	"os"
	"strings"

	"cloud.google.com/go/compute/metadata"
)

// Config holds all configuration for the application.
type Config struct {
	Service ServiceConfig
	Email   EmailConfig
	GCal    GCalConfig
	GCP     GCPConfig
	Ideas   IdeasConfig
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

	// More detailed check for missing environment variables
	var missingVars []string
	if smtpHost == "" {
		missingVars = append(missingVars, "SMTP_HOST")
	}
	if smtpPort == "" {
		missingVars = append(missingVars, "SMTP_PORT")
	}
	if sendFrom == "" {
		missingVars = append(missingVars, "SEND_FROM")
	}
	if smtpPass == "" {
		missingVars = append(missingVars, "SMTP_PASS")
	}
	if len(missingVars) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %s", strings.Join(missingVars, ", "))
	}

	calendarID := os.Getenv("CALENDAR_ID")
	availableSlotSummary := os.Getenv("GCAL_AVAILABLE_SLOT_SUMMARY")
	if calendarID == "" || availableSlotSummary == "" {
		return nil, fmt.Errorf("one or more required Google Calendar environment variables are not set (CALENDAR_ID, GCAL_AVAILABLE_SLOT_SUMMARY)")
	}

	// GCP Project and Location are needed for Vertex AI.
	// The preferred way to get the Project ID on Cloud Run is from the metadata server.
	projectID, err := metadata.ProjectID()
	if err != nil {
		// Fallback to env var if not on GCP or metadata server is unavailable.
		// This is useful for local development.
		projectID = os.Getenv("GCP_PROJECT_ID")
	}
	if projectID == "" {
		return nil, fmt.Errorf("GCP_PROJECT_ID could not be determined from metadata or environment variable")
	}

	location := os.Getenv("GCP_LOCATION")
	if location == "" {
		return nil, fmt.Errorf("GCP_LOCATION environment variable is not set")
	}

	generateIdeasPromptTemplate := os.Getenv("GENERATE_IDEAS_PROMPT_TEMPLATE")

	return &Config{
		Service: ServiceConfig{Port: port},
		Email:   EmailConfig{SmtpHost: smtpHost, SmtpPort: smtpPort, SendFrom: sendFrom, SendFromAlias: sendFromAlias, SmtpPass: smtpPass},
		GCal:    GCalConfig{CalendarID: calendarID, AvailableSlotSummary: availableSlotSummary},
		GCP:     GCPConfig{ProjectID: projectID, Location: location},
		Ideas:   IdeasConfig{GenerateIdeasPromptTemplate: generateIdeasPromptTemplate},
	}, nil
}
