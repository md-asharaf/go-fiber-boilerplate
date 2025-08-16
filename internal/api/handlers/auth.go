package handlers

import (
	"github.com/gofiber/fiber/v2"
	m "github.com/md-asharaf/go-fiber-boilerplate/internal/models"
	s "github.com/md-asharaf/go-fiber-boilerplate/internal/services"
	u "github.com/md-asharaf/go-fiber-boilerplate/internal/utils"
)

// AuthHandler handles authentication requests
type AuthHandler struct {
	authService *s.AuthService
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authService *s.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// Register handles user registration for Fiber
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var input m.RegisterInput
	if err := u.ParseAndValidateInput(c, &input); err != nil {
		return err
	}
	resp, err := h.authService.Register(input)
	if err != nil {
		return u.WriteErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	return u.WriteSuccessResponse(c, resp, "User registered successfully")
}

// Login handles user login for Fiber
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var input m.LoginInput
	if err := u.ParseAndValidateInput(c, &input); err != nil {
		return err
	}
	resp, err := h.authService.Login(input)
	if err != nil {
		return u.WriteErrorResponse(c, fiber.StatusUnauthorized, err.Error())
	}
	return u.WriteSuccessResponse(c, resp, "User logged in successfully")
}
