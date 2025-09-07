package analytics

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

const measurementProtocolURL = "https://www.google-analytics.com/mp/collect"

// Tracker sends events to the Google Analytics Measurement Protocol.
type Tracker struct {
	apiSecret     string
	measurementID string
	client        *http.Client
	logger        *slog.Logger
}

// NewTracker creates a new analytics tracker.
// It requires a GA4 Measurement Protocol API Secret and a Measurement ID.
func NewTracker(apiSecret, measurementID string, logger *slog.Logger) (*Tracker, error) {
	if apiSecret == "" {
		return nil, fmt.Errorf("analytics API secret is required")
	}
	if measurementID == "" {
		return nil, fmt.Errorf("analytics measurement ID is required")
	}
	return &Tracker{
		apiSecret:     apiSecret,
		measurementID: measurementID,
		client:        &http.Client{Timeout: 10 * time.Second},
		logger:        logger.With("service", "analytics_tracker"),
	}, nil
}

// BookingConfirmedEvent represents the data for a confirmed booking.
type BookingConfirmedEvent struct {
	ClientID      string
	TransactionID string
	Value         float64
	Currency      string
}

// TrackBookingConfirmed sends the highest-value conversion event to GA4.
func (t *Tracker) TrackBookingConfirmed(ctx context.Context, eventDetails BookingConfirmedEvent) {
	if eventDetails.ClientID == "" {
		t.logger.Warn("Cannot track booking confirmation: clientID is missing.")
		return
	}

	url := fmt.Sprintf("%s?api_secret=%s&measurement_id=%s", measurementProtocolURL, t.apiSecret, t.measurementID)

	payload := map[string]interface{}{
		"client_id": eventDetails.ClientID,
		"events": []map[string]interface{}{
			{
				"name": "booking_confirmed",
				"params": map[string]interface{}{
					"currency":             eventDetails.Currency,
					"value":                eventDetails.Value,
					"transaction_id":       eventDetails.TransactionID,
					"session_id":           "12345", // You can enhance this later if you capture session_id
					"engagement_time_msec": 1,
				},
			},
		},
	}

	body, err := json.Marshal(payload)
	if err != nil {
		t.logger.Error("Failed to marshal analytics payload", "error", err)
		return
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		t.logger.Error("Failed to create analytics request", "error", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := t.client.Do(req)
	if err != nil {
		t.logger.Error("Failed to send event to Google Analytics", "error", err)
		return
	}
	defer resp.Body.Close()

	// A 204 No Content response is a success.
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		t.logger.Info("Successfully sent 'booking_confirmed' event to Google Analytics", "client_id_suffix", eventDetails.ClientID[len(eventDetails.ClientID)-4:], "transaction_id", eventDetails.TransactionID)
	} else {
		t.logger.Error("Google Analytics API returned a non-success status", "status_code", resp.StatusCode)
	}
}
