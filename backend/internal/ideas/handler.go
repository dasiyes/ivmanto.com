package ideas

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/vertexai/genai"
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

// cleanJSONString removes markdown code blocks if they exist.
func cleanJSONString(s string) string {
	if strings.HasPrefix(s, "```json") {
		s = strings.TrimPrefix(s, "```json")
		s = strings.TrimSuffix(s, "```")
	}
	return strings.TrimSpace(s)
}

// Handler creates an HTTP handler for generating ideas.
func Handler(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, `{"error": "Only POST method is allowed"}`, http.StatusMethodNotAllowed)
			return
		}

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

		ctx := context.Background()
		projectID := os.Getenv("GCP_PROJECT_ID")
		location := os.Getenv("GCP_LOCATION")

		if projectID == "" || location == "" {
			logger.Error("GCP_PROJECT_ID or GCP_LOCATION environment variables not set.")
			http.Error(w, `{"error": "Server configuration error."}`, http.StatusInternalServerError)
			return
		}

		prompt := fmt.Sprintf(
			`You are a world-class data strategy consultant. A potential client has provided the following topic: '%s'. Generate 3 to 5 creative and insightful blog post titles based on this topic. For each title, provide a compelling one-sentence summary. Format the output as a valid JSON array of objects, where each object has a "title" and a "summary" field. Do not include any other text or explanations outside of the JSON array.`,
			req.Topic,
		)

		client, err := genai.NewClient(ctx, projectID, location)
		if err != nil {
			logger.Error("Error creating Vertex AI client", "error", err)
			http.Error(w, `{"error": "Could not connect to AI service."}`, http.StatusInternalServerError)
			return
		}
		defer client.Close()

		model := client.GenerativeModel("gemini-1.0-pro-001")

		resp, err := model.GenerateContent(ctx, genai.Text(prompt))
		if err != nil {
			logger.Error("Error generating content from Vertex AI", "error", err)
			http.Error(w, `{"error": "Failed to generate ideas from AI service."}`, http.StatusInternalServerError)
			return
		}

		if len(resp.Candidates) == 0 || resp.Candidates[0].Content == nil || len(resp.Candidates[0].Content.Parts) == 0 {
			logger.Error("AI service returned empty response")
			http.Error(w, `{"error": "AI service returned an empty response."}`, http.StatusInternalServerError)
			return
		}

		// The response part is of type genai.Text, which is an alias for string.
		aiResponseText := resp.Candidates[0].Content.Parts[0].(genai.Text)
		jsonStr := cleanJSONString(string(aiResponseText))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(jsonStr))
	}
}
