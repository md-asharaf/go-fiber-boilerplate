package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/go-backend-boilerplate/internal/models"
	"github.com/yourusername/go-backend-boilerplate/internal/services"
	"github.com/yourusername/go-backend-boilerplate/internal/utils"
)

// UserHandler handles user-related requests
type UserHandler struct {
	userService *services.UserService
}

// NewUserHandler creates a new user handler
func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Me returns the current authenticated user's info for Fiber
func (h *UserHandler) Me(c *fiber.Ctx) error {
	userValue := c.Locals("user")
	if userValue == nil {
		return utils.WriteErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized: missing user context")
	}
	user, ok := userValue.(*models.User)
	if !ok {
		return utils.WriteErrorResponse(c, fiber.StatusUnauthorized, "Unauthorized: invalid user context")
	}
	return utils.WriteSuccessResponse(c, user.ToResponse(), "User fetched successfully")
}
