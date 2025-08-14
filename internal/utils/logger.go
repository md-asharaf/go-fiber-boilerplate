package utils

import (
	"os"

	"github.com/md-asharaf/go-fiber-boilerplate/internal/config"
	"go.uber.org/zap"
)

// Global logger instance
var Logger *zap.Logger

// InitLogger initializes the global Logger variable
func InitLogger(cfg config.LoggingConfig) error {
	var zapConfig zap.Config

	// Check environment to determine log format
	env := os.Getenv("ENV")
	if env == "production" {
		// Use production config (JSON format, structured)
		zapConfig = zap.NewProductionConfig()
	} else {
		// Use development config (console format, human-readable)
		zapConfig = zap.NewDevelopmentConfig()
	}

	// Set log level
	level, err := zap.ParseAtomicLevel(cfg.Level)
	if err != nil {
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}
	zapConfig.Level = level

	// Build logger
	logger, err := zapConfig.Build(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel),
	)
	if err != nil {
		return err
	}
	Logger = logger
	return nil
}
