package cmd

import (
	"encoding/json"
	"net/http"
	"time"
)

// StandardResponse 구조체
type StandardResponse struct {
	Code      int       `json:"code"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	// Data    interface{} `json:"data,omitempty"`
	Data json.RawMessage `json:"data,omitempty"`
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Respond with a standard not found message
	response := StandardResponse{
		Code:      http.StatusNotFound,
		Message:   "Not Found",
		Timestamp: time.Now(),
		Data:      nil,
	}

	// Encode the response as JSON and write it to the response writer
	w.WriteHeader(http.StatusNotFound)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	// Create a sample JSON response
	response := StandardResponse{
		Code:      http.StatusOK,
		Message:   "Success",
		Timestamp: time.Now(),
		Data:      nil,
	}
	// Encode the response as JSON and write it to the response writer
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}
