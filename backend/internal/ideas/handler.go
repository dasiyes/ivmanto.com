package ideas

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"cloud.google.com/go/vertexai/genai"
	"ivmanto.com/backend/internal/email"
)

// ModelName is the specific Vertex AI model to use for generating ideas.
// We use gemini-1.5-pro as it is available in the europe-west3 region.
// The original gemini-1.0-pro is not available there, which caused the error.
const ModelName = "gemini-1.5-pro"

// Handler manages ideas-related HTTP requests.
type Handler struct {
	logger      *slog.Logger
	genaiClient *genai.Client
	emailSvc    email.Service
}

// NewHandler creates a new ideas handler.
func NewHandler(logger *slog.Logger, genaiClient *genai.Client, emailSvc email.Service) *Handler {
	return &Handler{
		logger:      logger,
		genaiClient: genaiClient,
		emailSvc:    emailSvc,
	}
}

// RegisterRoutes sets up the routing for ideas endpoints.
func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/generate-ideas", h.handleGenerateIdeas)
	mux.HandleFunc("POST /api/ideas/email", h.handleEmailIdeas)
}

// --- Response Helpers ---
func (h *Handler) respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if payload != nil {
		if err := json.NewEncoder(w).Encode(payload); err != nil {
			h.logger.Error("could not write JSON response", "error", err)
		}
	}
}

func (h *Handler) respondError(w http.ResponseWriter, status int, message string) {
	h.respondJSON(w, status, map[string]string{"error": message})
}

// --- Generate Ideas Handler Logic ---

// GenerateIdeasRequest represents the expected JSON body for a generate-ideas request.
type GenerateIdeasRequest struct {
	Topic string `json:"topic"`
}

// GenerateIdeasResponse represents the JSON response containing the generated ideas.
type GenerateIdeasResponse struct {
	Ideas []string `json:"ideas"`
}

func (h *Handler) handleGenerateIdeas(w http.ResponseWriter, r *http.Request) {
	var req GenerateIdeasRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	if req.Topic == "" {
		h.respondError(w, http.StatusBadRequest, "Topic cannot be empty")
		return
	}

	h.logger.Info("Received idea generation request", "topic", req.Topic)

	model := h.genaiClient.GenerativeModel(ModelName)
	model.SetTemperature(0.8)

	prompt := genai.Text("Generate a numbered list of 5 creative and actionable business ideas related to: " + req.Topic)

	resp, err := model.GenerateContent(r.Context(), prompt)
	if err != nil {
		h.logger.Error("Error generating content from Vertex AI", "error", err, "model", ModelName)
		h.respondError(w, http.StatusInternalServerError, "Failed to generate ideas from AI model")
		return
	}

	var generatedText string
	if len(resp.Candidates) > 0 && len(resp.Candidates[0].Content.Parts) > 0 {
		if txt, ok := resp.Candidates[0].Content.Parts[0].(genai.Text); ok {
			generatedText = string(txt)
		}
	}

	if generatedText == "" {
		h.logger.Warn("Vertex AI returned a response with no text content", "topic", req.Topic)
		h.respondError(w, http.StatusInternalServerError, "AI model returned an empty response")
		return
	}

	h.respondJSON(w, http.StatusOK, GenerateIdeasResponse{Ideas: strings.Split(generatedText, "\n")})
}

// --- Email Ideas Handler Logic ---

// EmailIdeasRequest is the structure for the email ideas request.
type EmailIdeasRequest struct {
	Email string   `json:"email"`
	Topic string   `json:"topic"`
	Ideas []string `json:"ideas"`
}

func (h *Handler) handleEmailIdeas(w http.ResponseWriter, r *http.Request) {
	var req EmailIdeasRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("Failed to decode email ideas request", "error", err)
		h.respondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Email == "" || req.Topic == "" || len(req.Ideas) == 0 {
		h.respondError(w, http.StatusBadRequest, "Email, topic, and ideas are required")
		return
	}

	h.logger.Info("Received request to email ideas", "email", req.Email, "topic", req.Topic)

	// The email service likely expects a single string for the body, not a slice.
	// We'll format the ideas into a single string here.
	emailBody := strings.Join(req.Ideas, "\n")

	err := h.emailSvc.SendGeneratedIdeas(req.Email, req.Topic, emailBody)
	if err != nil {
		h.logger.Error("Failed to send generated ideas email", "error", err)
		h.respondError(w, http.StatusInternalServerError, "Failed to send email")
		return
	}

	h.respondJSON(w, http.StatusOK, map[string]string{"message": "Email sent successfully"})
}
