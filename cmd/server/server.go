package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/config"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/utils"
	"go.uber.org/zap"
)

func StartServer(app *fiber.App, cfg config.ServerConfig) {
	serverAddr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	utils.Logger.Info("Starting server", zap.String("address", serverAddr))

	if err := app.Listen(serverAddr); err != nil {
		utils.Logger.Fatal("Failed to start server", zap.Error(err))
	}
}
