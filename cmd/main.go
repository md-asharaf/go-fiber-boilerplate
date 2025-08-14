package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/md-asharaf/go-fiber-boilerplate/cmd/server"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/api/routes"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/config"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/utils"
	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize services
	appServices, err := services.InitServices(cfg)
	if err != nil {
		utils.Logger.Error("failed to initialize services", zap.Error(err))
		return
	}

	// Setup Fiber app
	app := fiber.New()
	// Pass services to routes for global access
	routes.SetupRoutes(app, appServices)

	server.StartServer(app, cfg.Server)

	// On shutdown, close services and log any errors
	if err := appServices.CloseServices(); err != nil {
		utils.Logger.Error("failed to close services", zap.Error(err))
	}
}
