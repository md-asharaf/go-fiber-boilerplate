package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/api/handlers/auth"
)

func CreateAuthRoutes(api fiber.Router, userHandler *auth.AuthHandler) {
	protected := api.Group("/auth")
	protected.Post("/register", userHandler.Register)
	protected.Post("/login", userHandler.Login)
}
