package routes

import (
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/yourusername/go-backend-boilerplate/internal/api/handlers"
	"github.com/yourusername/go-backend-boilerplate/internal/api/middleware"
	"github.com/yourusername/go-backend-boilerplate/internal/services"
)

// SetupRoutes configures all application routes
func SetupRoutes(router chi.Router, appServices *services.AppServices) chi.Router {
	// Global middleware
	router.Use(chimiddleware.Logger)
	router.Use(chimiddleware.Recoverer)
	router.Use(chimiddleware.StripSlashes)
	router.Use(middleware.CORS())

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(appServices.AuthService)
	healthHandler := handlers.NewHealthHandler()
	userHandler := handlers.NewUserHandler(appServices.UserService)

	// API routes
	router.Route("/api/v1", func(r chi.Router) {
		// Health check (no auth required)
		r.Get("/health", healthHandler.HealthCheck)

		// Auth routes (no auth required)
		r.Post("/auth/register", authHandler.Register)
		r.Post("/auth/login", authHandler.Login)

		// Protected routes with authentication
		r.Group(func(r chi.Router) {
			r.Use(middleware.JWTAuth(appServices.JWTService, appServices.UserService))
			r.Get("/user/me", userHandler.Me)
		})
	})

	return router
}
