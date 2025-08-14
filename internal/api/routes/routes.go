package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yourusername/go-backend-boilerplate/internal/api/handlers"
	"github.com/yourusername/go-backend-boilerplate/internal/api/middleware"
	"github.com/yourusername/go-backend-boilerplate/internal/services"
)

// SetupFiberRoutes configures all application routes for Fiber
func SetupFiberRoutes(app *fiber.App, appServices *services.AppServices) {
	// Global middleware
	app.Use(middleware.CORS())

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(appServices.AuthService)
	userHandler := handlers.NewUserHandler(appServices.UserService)

	api := app.Group("/api/v1")

	// Health check (no auth required)
	api.Get("/health", handlers.HealthCheck)

	// Auth routes (no auth required)
	api.Post("/auth/register", authHandler.Register)
	api.Post("/auth/login", authHandler.Login)

	// Protected routes with authentication
	protected := api.Group("/", middleware.JWTAuth(appServices.JWTService, appServices.UserService))
	protected.Get("/user/me", userHandler.Me)
}
