package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/models"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/utils"
)

// AuthHandler handles authentication requests
type AuthHandler struct {
	authService *services.AuthService
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register handles user registration for Fiber
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var input models.RegisterInput
	if err := c.BodyParser(&input); err != nil {
		return utils.WriteErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}
	if err := utils.ValidateStruct(&input); err != nil {
		return utils.WriteErrorResponse(c, fiber.StatusBadRequest, "Validation failed: "+err.Error())
	}
	resp, err := h.authService.Register(input)
	if err != nil {
		return utils.WriteErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	return utils.WriteSuccessResponse(c, resp, "User registered successfully")
}

// Login handles user login for Fiber
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var input models.LoginInput
	if err := c.BodyParser(&input); err != nil {
		return utils.WriteErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}
	if err := utils.ValidateStruct(&input); err != nil {
		return utils.WriteErrorResponse(c, fiber.StatusBadRequest, "Validation failed: "+err.Error())
	}
	resp, err := h.authService.Login(input)
	if err != nil {
		return utils.WriteErrorResponse(c, fiber.StatusUnauthorized, err.Error())
	}
	return utils.WriteSuccessResponse(c, resp, "User logged in successfully")
}
