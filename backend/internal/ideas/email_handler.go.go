package ideas

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"ivmanto.com/backend/internal/email"
)

// EmailIdeasRequest is the structure for the email ideas request.
type EmailIdeasRequest struct {
	Email string         `json:"email"`
	Topic string         `json:"topic"`
	Ideas []IdeaResponse `json:"ideas"` // Re-using IdeaResponse from handler.go
}

// EmailHandler creates an HTTP handler for emailing generated ideas.
func EmailHandler(logger *slog.Logger, emailSvc email.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, `{"error": "Only POST method is allowed"}`, http.StatusMethodNotAllowed)
			return
		}

		var req EmailIdeasRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			logger.Error("Failed to decode email ideas request", "error", err)
			http.Error(w, `{"error": "Invalid request body"}`, http.StatusBadRequest)
			return
		}

		if req.Email == "" || req.Topic == "" || len(req.Ideas) == 0 {
			http.Error(w, `{"error": "Email, topic, and ideas are required"}`, http.StatusBadRequest)
			return
		}

		logger.Info("Received request to email ideas", "email", req.Email, "topic", req.Topic)

		ideasForEmail := make([]email.GeneratedIdea, len(req.Ideas))
		for i, idea := range req.Ideas {
			ideasForEmail[i] = email.GeneratedIdea(idea)
		}

		err := emailSvc.SendGeneratedIdeas(req.Email, req.Topic, ideasForEmail)
		if err != nil {
			logger.Error("Failed to send generated ideas email", "error", err)
			http.Error(w, `{"error": "Failed to send email"}`, http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Email sent successfully"})
	}
}
