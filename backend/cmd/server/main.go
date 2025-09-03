package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"ivmanto.com/backend/internal/booking"
	"ivmanto.com/backend/internal/config"
	"ivmanto.com/backend/internal/contact"
	"ivmanto.com/backend/internal/email"
	"ivmanto.com/backend/internal/gcal"
	"ivmanto.com/backend/internal/ideas"
	"ivmanto.com/backend/internal/middleware"
)

func main() {
	// Load .env file for local development.
	err := godotenv.Load()
	if err != nil {
		slog.Info(".env file not found, loading config from environment")
	}

	// 1. Initialize structured logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger) // Set as default for convenience

	// 2. Load configuration
	cfg, err := config.Load()
	if err != nil {
		slog.Error("could not load config", "error", err)
		os.Exit(1)
	}

	// 3. Initialize services
	emailService := email.NewSmtpService(&cfg.Email)
	ctx := context.Background()
	gcpCredsPath := "gcp-credentials.json"
	if prodPath := os.Getenv("GCP_CREDENTIALS_PATH"); prodPath != "" {
		gcpCredsPath = prodPath
	}

	gcalSvc, err := gcal.NewService(ctx, cfg, gcpCredsPath)
	if err != nil {
		slog.Error("Failed to create Google Calendar service", "error", err)
		os.Exit(1)
	}

	// 4. Initialize handlers, passing dependencies
	contactHandler := contact.NewHandler(logger, emailService)
	bookingHandler := booking.NewHandler(logger, gcalSvc, emailService)
	ideasGenerateHandler := ideas.Handler(logger)
	ideasEmailHandler := ideas.EmailHandler(logger, emailService)

	// 5. Register routes
	mux := http.NewServeMux()
	contactHandler.RegisterRoutes(mux)
	bookingHandler.RegisterRoutes(mux)
	mux.HandleFunc("POST /api/generate-ideas", ideasGenerateHandler)
	mux.HandleFunc("POST /api/ideas/email", ideasEmailHandler)

	// 6. Apply middleware
	var finalHandler http.Handler = mux
	finalHandler = middleware.Cors(finalHandler)
	finalHandler = middleware.RequestLogger(logger, finalHandler)

	// 7. Start server
	slog.Info("Starting server", "port", cfg.Service.Port)
	server := &http.Server{
		Addr:    ":" + cfg.Service.Port,
		Handler: finalHandler,
	}
	if err := server.ListenAndServe(); err != nil {
		slog.Error("could not start server", "error", err)
		os.Exit(1)
	}
}
