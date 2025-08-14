package utils

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ErrorResponse represents a standard error response for Fiber
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// SuccessResponse represents a standard success response for Fiber
type SuccessResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
}

// WriteErrorResponse writes a standard error response in Fiber
func WriteErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
		Status:  statusCode,
	})
}

// WriteSuccessResponse writes a standard success response in Fiber
func WriteSuccessResponse(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Data:    data,
		Message: message,
		Status:  "success",
	})
}
