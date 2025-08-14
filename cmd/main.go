package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/api/routes"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/config"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/database"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/models"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/utils"
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

	// Setup Fiber app
	app := fiber.New()

	// Pass appServices to routes for global access
	routes.SetupRoutes(app, appServices)

	// Start server
	serverAddr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	zap.L().Info("Starting server", zap.String("address", serverAddr))

	if err := app.Listen(serverAddr); err != nil {
		zap.L().Fatal("Failed to start server", zap.Error(err))
	}
}
