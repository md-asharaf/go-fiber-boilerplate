package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/api/handlers"
	authHandler "github.com/md-asharaf/go-fiber-boilerplate/internal/api/handlers/auth"
	userHandler "github.com/md-asharaf/go-fiber-boilerplate/internal/api/handlers/user"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/api/middleware"
	authRoutes "github.com/md-asharaf/go-fiber-boilerplate/internal/api/routes/auth"
	userRoutes "github.com/md-asharaf/go-fiber-boilerplate/internal/api/routes/user"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services/auth"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services/email"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services/jwt"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services/otp"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services/redis"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services/user"
)

type Services struct {
	AuthService  *auth.AuthService
	UserService  *user.UserService
	JWTService   *jwt.JWTService
	RedisService *redis.RedisService
	EmailService *email.EmailService
	OtpService   *otp.OtpService
}

// SetupFiberRoutes configures all application routes for Fiber
func SetupRoutes(app *fiber.App, services *Services) {
	// Global middleware
	app.Use(middleware.CORS())

	// Initialize handlers
	authHandler := authHandler.NewAuthHandler(services.AuthService)
	userHandler := userHandler.NewUserHandler(services.UserService)

	authMiddleware := middleware.JWTAuth(services.JWTService, services.UserService)

	api := app.Group("/api/v1")
	// Health check (no auth required)
	api.Get("/health", handlers.HealthCheck)

	authRoutes.CreateAuthRoutes(api, authHandler)
	userRoutes.CreateUserRoutes(api, userHandler, authMiddleware)
}
