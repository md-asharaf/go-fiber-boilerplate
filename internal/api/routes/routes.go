package routes

import (
	"github.com/gofiber/fiber/v2"
	h "github.com/md-asharaf/go-fiber-boilerplate/internal/api/handlers"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/api/middleware"
	s "github.com/md-asharaf/go-fiber-boilerplate/internal/services"
)

type Services struct {
	AuthService  *s.AuthService
	UserService  *s.UserService
	JWTService   *s.JWTService
	RedisService *s.RedisService
	EmailService *s.EmailService
	OtpService   *s.OtpService
}

// SetupFiberRoutes configures all application routes for Fiber
func SetupRoutes(app *fiber.App, services *Services) {
	// Global middleware
	app.Use(middleware.CORS())

	// Initialize handlers
	authHandler := h.NewAuthHandler(services.AuthService)
	userHandler := h.NewUserHandler(services.UserService)

	authMiddleware := middleware.JWTAuth(services.JWTService, services.UserService)

	api := app.Group("/api/v1")
	// Health check (no auth required)
	api.Get("/health", h.HealthCheck)

	CreateAuthRoutes(api, authHandler)
	CreateUserRoutes(api, userHandler, authMiddleware)
}
