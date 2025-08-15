package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/config"
	"go.uber.org/zap"
)

func StartServer(app *fiber.App, cfg config.ServerConfig, logger *zap.Logger) {
	serverAddr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	logger.Info("Starting server", zap.String("address", serverAddr))

	if err := app.Listen(serverAddr); err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
