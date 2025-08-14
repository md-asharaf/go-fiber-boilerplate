package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/yourusername/go-backend-boilerplate/internal/api/routes"
	"github.com/yourusername/go-backend-boilerplate/internal/config"
	"github.com/yourusername/go-backend-boilerplate/internal/database"
	"github.com/yourusername/go-backend-boilerplate/internal/models"
	"github.com/yourusername/go-backend-boilerplate/internal/services"
	"github.com/yourusername/go-backend-boilerplate/internal/utils"
	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Setup logging
	utils.SetupLogging(cfg.Logging)

	zap.L().Info("Starting application...")

	// Connect to database
	if err := database.Connect(cfg.Database); err != nil {
		zap.L().Fatal("Failed to connect to database", zap.Error(err))
	}

	// Get database instance
	db := database.GetDB()

	// Auto-migrate database
	if err := db.AutoMigrate(&models.User{}); err != nil {
		zap.L().Fatal("Failed to migrate database", zap.Error(err))
	}

	// Initialize services
	appServices := services.InitServices(cfg, db)

	// Setup router
	router := chi.NewRouter()

	// Pass appServices to routes for global access
	routes.SetupRoutes(router, appServices)

	// Start server
	serverAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	zap.L().Info("Starting server", zap.String("address", serverAddr))

	err := http.ListenAndServe(serverAddr, router)
	if err != nil {
		zap.L().Fatal("Failed to start server", zap.Error(err))
	}
}
