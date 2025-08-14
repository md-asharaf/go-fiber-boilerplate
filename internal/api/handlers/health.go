package handlers

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

// HealthHandler handles health-related HTTP requests
type HealthHandler struct{}

// NewHealthHandler creates a new health handler
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck handles GET /health request
func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	zap.L().Info("Health check requested")

	response := map[string]interface{}{
		"status":    "OK",
		"service":   "go-crypto-api",
		"timestamp": "2025-08-14T13:31:55+05:30",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		zap.L().Error("Failed to encode health response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
