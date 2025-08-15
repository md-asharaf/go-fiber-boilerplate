package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/models"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services/auth"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/utils"
)

// AuthHandler handles authentication requests
type AuthHandler struct {
	authService *auth.AuthService
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService *auth.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register handles user registration for Fiber
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var input models.RegisterInput
	if err := utils.ParseAndValidateInput(c, &input); err != nil {
		return err
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
	if err := utils.ParseAndValidateInput(c, &input); err != nil {
		return err
	}
	resp, err := h.authService.Login(input)
	if err != nil {
		return utils.WriteErrorResponse(c, fiber.StatusUnauthorized, err.Error())
	}
	return utils.WriteSuccessResponse(c, resp, "User logged in successfully")
}
