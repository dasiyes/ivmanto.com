package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"ivmanto.com/backend/internal/booking"
	"ivmanto.com/backend/internal/config"
	"ivmanto.com/backend/internal/contact"
	"ivmanto.com/backend/internal/email"
	"ivmanto.com/backend/internal/gcal"
)

func main() {
	// Load .env file for local development.
	// In production (like Cloud Run), environment variables are set directly.
	err := godotenv.Load()
	if err != nil {
		// We don't want to fail if the .env file is missing,
		// as it's optional for production environments.
		log.Println("INFO: .env file not found, loading config from environment")
	}

	// 1. Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("FATAL: could not load config: %s\n", err)
	}

	// 2. Initialize services
	// The email service is kept for the contact form and future notifications.
	emailService := email.NewSmtpService(&cfg.Email)

	// Define the path for the service account credentials.
	// For local dev, we point to the file in the backend dir.
	// For production, this path will be mounted by Cloud Run from Secret Manager.
	gcpCredsPath := "gcp-credentials.json"
	if prodPath := os.Getenv("GCP_CREDENTIALS_PATH"); prodPath != "" {
		gcpCredsPath = prodPath
	}

	// Initialize the Google Calendar service. This is the core of our new booking engine.
	ctx := context.Background()
	// We now pass the path to the credentials file to enable Domain-Wide Delegation.
	gcalSvc, err := gcal.NewService(ctx, cfg, gcpCredsPath)
	if err != nil {
		log.Fatalf("FATAL: Failed to create Google Calendar service: %v", err)
	}

	// 3. Initialize handlers.
	contactHandler := contact.NewHandler(emailService)
	bookingHandler := booking.NewHandler(gcalSvc, emailService)

	// 4. Register routes
	mux := http.NewServeMux()
	contactHandler.RegisterRoutes(mux)
	bookingHandler.RegisterRoutes(mux)

	log.Printf("INFO: Starting server on :%s", cfg.Service.Port)
	if err := http.ListenAndServe(":"+cfg.Service.Port, mux); err != nil {
		log.Fatalf("FATAL: could not start server: %s\n", err)
	}
}
