package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"cloud.google.com/go/vertexai/genai"
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

	// Get the port from the environment variable.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
		slog.Info("Defaulting to port", "number", port)
	}

	// 2. Load configuration
	cfg, err := config.Load()
	if err != nil {
		slog.Error("could not load config", "error", err)
		os.Exit(1)
	}

	// 3. Initialize services
	emailService := email.NewSmtpService(&cfg.Email, logger)
	ctx := context.Background()

	// For GCal, we use Application Default Credentials (ADC).
	// On Cloud Run, this uses the attached service account's identity.
	// For local development, run `gcloud auth application-default login`.
	gcpCredsPath := "" // An empty path forces the client library to use ADC.
	gcalSvc, err := gcal.NewService(ctx, cfg, gcpCredsPath)
	if err != nil {
		slog.Error("Failed to create Google Calendar service", "error", err)
		os.Exit(1)
	}

	// Initialize Vertex AI client once.
	genaiClient, err := genai.NewClient(ctx, cfg.GCP.ProjectID, cfg.GCP.Location)
	if err != nil {
		slog.Error("Failed to create Vertex AI client", "error", err)
		os.Exit(1)
	}
	defer genaiClient.Close()

	// 4. Initialize handlers, passing dependencies
	contactHandler := contact.NewHandler(logger, emailService)
	bookingHandler := booking.NewHandler(logger, gcalSvc, emailService)
	ideasHandler := ideas.NewHandler(logger, genaiClient, emailService)

	// 5. Register routes
	mux := http.NewServeMux()
	contactHandler.RegisterRoutes(mux)
	bookingHandler.RegisterRoutes(mux)
	ideasHandler.RegisterRoutes(mux)

	// 6. Apply middleware
	var finalHandler http.Handler = mux
	finalHandler = middleware.Cors(finalHandler)
	finalHandler = middleware.RequestLogger(logger, finalHandler)

	// 7. Start server
	slog.Info("Starting server", "port", port)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: finalHandler,
	}
	if err := server.ListenAndServe(); err != nil {
		slog.Error("could not start server", "error", err)
		os.Exit(1)
	}
}
