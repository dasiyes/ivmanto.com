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

	// Read the GCP credentials file
	creds, err := os.ReadFile("gcp-credentials.json")
	if err != nil {
		log.Fatalf("FATAL: Failed to read gcp-credentials.json: %v", err)
	}

	// Initialize the Google Calendar service. This is the core of our new booking engine.
	ctx := context.Background()
	gcalSvc, err := gcal.NewService( // Assuming gcal.NewService is updated to accept credentials as a string/byte slice
		ctx,
		string(creds),
		cfg.Booking.CalendarID,
		cfg.Booking.AvailableSlotSummary,
	)
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
