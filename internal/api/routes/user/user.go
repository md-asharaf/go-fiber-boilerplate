package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/api/handlers/user"
)

func CreateUserRoutes(api fiber.Router, userHandler *user.UserHandler, middleware fiber.Handler) {
	protected := api.Group("/user", middleware)
	protected.Get("/me", userHandler.Me)
}
