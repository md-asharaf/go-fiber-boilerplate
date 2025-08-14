package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/api/handlers"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/api/middleware"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services"
)

// SetupFiberRoutes configures all application routes for Fiber
func SetupRoutes(app *fiber.App, services *services.Services) {
	// Global middleware
	app.Use(middleware.CORS())

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(services.AUTH)
	userHandler := handlers.NewUserHandler(services.USER)

	api := app.Group("/api/v1")

	// Health check (no auth required)
	api.Get("/health", handlers.HealthCheck)

	// Auth routes (no auth required)
	api.Post("/auth/register", authHandler.Register)
	api.Post("/auth/login", authHandler.Login)

	// Protected routes with authentication
	protected := api.Group("/", middleware.JWTAuth(services.JWT, services.USER))
	protected.Get("/user/me", userHandler.Me)
}
