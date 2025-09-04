package ideas

import (
	"encoding/json"
	"fmt"
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
	logger         *slog.Logger
	genaiClient    *genai.Client
	emailSvc       email.Service
	promptTemplate string
}

// NewHandler creates a new ideas handler.
func NewHandler(logger *slog.Logger, genaiClient *genai.Client, emailSvc email.Service, promptTemplate string) *Handler {
	// Provide a robust default if the prompt isn't configured via environment variables.
	if promptTemplate == "" {
		logger.Warn("Generate ideas prompt template is not configured, using default.")
		promptTemplate = "You are a world-class data strategy consultant. A potential client has provided the following topic: '%s'. Generate 3 to 5 creative and insightful blog post titles based on this topic. For each title, provide a compelling one-sentence summary. Respond with ONLY a valid JSON array of objects. Each object must have a 'title' (string) and a 'summary' (string) field. Do not include any other text or markdown formatting outside of the JSON array."
	}

	return &Handler{
		logger:         logger,
		genaiClient:    genaiClient,
		emailSvc:       emailSvc,
		promptTemplate: promptTemplate,
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

// Idea represents a single generated idea with a title and summary.
type Idea struct {
	Title   string `json:"title"`
	Summary string `json:"summary"`
}

// GenerateIdeasResponse represents the JSON response containing the generated ideas.
type GenerateIdeasResponse struct {
	Ideas []Idea `json:"ideas"`
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

	// Use the configured prompt template.
	promptText := fmt.Sprintf(h.promptTemplate, req.Topic)
	prompt := genai.Text(promptText)

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

	h.logger.Info("Raw response from Vertex AI", "text", generatedText)

	// Clean the response: LLMs sometimes wrap JSON in ```json ... ```
	cleanedJSON := strings.TrimSpace(generatedText)
	cleanedJSON = strings.TrimPrefix(cleanedJSON, "```json")
	cleanedJSON = strings.TrimPrefix(cleanedJSON, "```")
	cleanedJSON = strings.TrimSuffix(cleanedJSON, "```")
	cleanedJSON = strings.TrimSpace(cleanedJSON)

	var ideas []Idea
	if err := json.Unmarshal([]byte(cleanedJSON), &ideas); err != nil {
		h.logger.Error("Failed to unmarshal JSON from Vertex AI", "error", err, "raw_text", generatedText)
		h.respondError(w, http.StatusInternalServerError, "AI model returned an invalid format.")
		return
	}

	if len(ideas) == 0 {
		h.logger.Warn("AI response resulted in zero ideas after parsing", "raw_text", generatedText)
		h.respondError(w, http.StatusInternalServerError, "AI model returned an empty or invalid response")
		return
	}

	h.respondJSON(w, http.StatusOK, GenerateIdeasResponse{Ideas: ideas})
}

// --- Email Ideas Handler Logic ---

// EmailIdeasRequest is the structure for the email ideas request.
type EmailIdeasRequest struct {
	Email string `json:"email"`
	Topic string `json:"topic"`
	Ideas []Idea `json:"ideas"`
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

	// Format the ideas into a single HTML string for the email body.
	var bodyBuilder strings.Builder
	bodyBuilder.WriteString("<ul>")
	for _, idea := range req.Ideas {
		bodyBuilder.WriteString(fmt.Sprintf("<li><strong>%s</strong><br>%s</li>", idea.Title, idea.Summary))
	}
	bodyBuilder.WriteString("</ul>")
	emailBody := bodyBuilder.String()

	err := h.emailSvc.SendGeneratedIdeas(req.Email, req.Topic, emailBody)
	if err != nil {
		h.logger.Error("Failed to send generated ideas email", "error", err)
		h.respondError(w, http.StatusInternalServerError, "Failed to send email")
		return
	}

	h.respondJSON(w, http.StatusOK, map[string]string{"message": "Email sent successfully"})
}
