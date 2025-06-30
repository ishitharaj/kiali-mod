package handlers

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"strings"

	"github.com/kiali/kiali/business"
	"github.com/kiali/kiali/log"
)

// LLMResponse represents the JSON response from the LLM
type LLMResponse struct {
	Answer string `json:"answer"`
}

// LLMHandler handles the LLM assistant requests
func LLMHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
		return
	}

	// Fetch metrics, validations, and service graph data
	graphData := "[REDACTED]" // Placeholder for actual graph data fetching logic
	metricsData := "[REDACTED]" // Placeholder for actual metrics data fetching logic

	// Sanitize and construct the prompt
	prompt := sanitizeAndConstructPrompt(query, graphData, metricsData)

	// Run the LLM locally
	cmd := exec.Command("ollama", "run", "llama3", prompt)
	output, err := cmd.Output()
	if err != nil {
		log.Errorf("Error running LLM: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Return the response as JSON
	response := LLMResponse{Answer: string(output)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// sanitizeAndConstructPrompt sanitizes sensitive data and constructs the LLM prompt
func sanitizeAndConstructPrompt(question, graphData, metricsData string) string {
	// Sanitize sensitive information
	question = strings.ReplaceAll(question, "sensitive", "[REDACTED]")
	// Construct the prompt
	prompt := "You are an observability assistant. Here's data:\n\n"
	prompt += "Graph: " + graphData + "\n"
	prompt += "Metrics: " + metricsData + "\n"
	prompt += "Question: " + question + "\n\n"
	prompt += "Respond with insights and potential root causes in plain language."
	return prompt
} 