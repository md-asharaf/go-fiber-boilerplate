package utils

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

// ErrorResponse represents a standard error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// SuccessResponse represents a standard success response
type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}

// WriteJSONResponse writes a JSON response to the HTTP response writer
func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		zap.L().Error("Failed to encode JSON response", zap.Error(err))
		return err
	}
	return nil
}

// WriteErrorResponse writes a standard error response
func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	errorResp := ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
		Status:  statusCode,
	}

	if err := WriteJSONResponse(w, statusCode, errorResp); err != nil {
		// Fallback to plain text if JSON encoding fails
		http.Error(w, message, statusCode)
	}
}

// WriteSuccessResponse writes a standard success response
func WriteSuccessResponse(w http.ResponseWriter, data interface{}, message string) {
	successResp := SuccessResponse{
		Data:    data,
		Message: message,
		Status:  "success",
	}

	if err := WriteJSONResponse(w, http.StatusOK, successResp); err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "Failed to encode response")
	}
}
