package routes

import (
	"github.com/gofiber/fiber/v2"
	h "github.com/md-asharaf/go-fiber-boilerplate/internal/api/handlers"
)

func CreateUserRoutes(api fiber.Router, userHandler *h.UserHandler, middleware fiber.Handler) {
	protected := api.Group("/user", middleware)
	protected.Get("/me", userHandler.Me)
}
