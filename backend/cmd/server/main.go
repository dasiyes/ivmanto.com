package main

import (
	"log"
	"net/http"

	"ivmanto.com/backend/internal/booking"
	"ivmanto.com/backend/internal/config"
	"ivmanto.com/backend/internal/email"
)

func main() {
	// 1. Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("could not load config: %s\n", err)
	}

	// 2. Initialize dependencies (storage, emailer, etc.).
	bookingStore := booking.NewInMemoryBookingStore()

	// Use the MockService for local dev to avoid needing real SMTP credentials.
	// To use the real SMTP service, comment out the MockService and uncomment the SmtpService.
	// the SMTP password is loaded into the config (e.g., from an environment variable).
	var emailService email.Service = email.NewMockService()
	// emailService := email.NewSmtpService(&cfg.Email)

	// 3. Initialize services and handlers
	bookingService := booking.NewBookingService(bookingStore, emailService, &cfg.Booking)
	bookingHandler := booking.NewHandler(bookingService)

	mux := http.NewServeMux()
	bookingHandler.RegisterRoutes(mux)

	// TODO: Add contact form handler registration here later.

	log.Printf("Starting server on :%s", cfg.Service.Port)
	if err := http.ListenAndServe(":"+cfg.Service.Port, mux); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
