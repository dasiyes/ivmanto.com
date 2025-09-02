package ideas

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

// GenerateIdeasRequest is the expected structure of the request body.
type GenerateIdeasRequest struct {
	Topic string `json:"topic"`
}

// IdeaResponse defines the structure for a single generated idea.
type IdeaResponse struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
}

// Handler creates an HTTP handler for generating ideas.
// For now, it returns mock data to test the frontend-backend connection.
func Handler(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req GenerateIdeasRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
			return
		}

		if req.Topic == "" {
			http.Error(w, `{"error": "Topic cannot be empty"}`, http.StatusBadRequest)
			return
		}

		logger.Info("Received topic for idea generation", "topic", req.Topic)

		// --- Placeholder for Gemini API call ---
		// Simulate network delay for the AI call
		time.Sleep(2 * time.Second)

		// Mock response
		mockIdeas := []IdeaResponse{
			{
				Title:   "Unlocking Retail's Future: How AI is Revolutionizing Customer Experience",
				Summary: "Explore five practical ways artificial intelligence is personalizing shopping, optimizing supply chains, and driving sales in the retail sector.",
			},
			{
				Title:   "Beyond the Hype: A Realistic Look at Implementing AI in Retail",
				Summary: "This article cuts through the noise to provide a step-by-step guide for retail businesses looking to adopt AI, from data readiness to measuring ROI.",
			},
			{
				Title:   "The Ethical Checkout: Navigating AI's Role in Retail Privacy and Trust",
				Summary: "As AI becomes more integrated into retail, we discuss the critical importance of building customer trust through transparent and ethical data practices.",
			},
		}
		// --- End Placeholder ---

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(mockIdeas)
	}
}
