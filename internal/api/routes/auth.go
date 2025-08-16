package routes

import (
	"github.com/gofiber/fiber/v2"
	h "github.com/md-asharaf/go-fiber-boilerplate/internal/api/handlers"
)

func CreateAuthRoutes(api fiber.Router, userHandler *h.AuthHandler) {
	protected := api.Group("/auth")
	protected.Post("/register", userHandler.Register)
	protected.Post("/login", userHandler.Login)
}
