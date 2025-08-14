package utils

import (
	"os"

	"github.com/yourusername/go-backend-boilerplate/internal/config"
	"go.uber.org/zap"
)

// SetupLogging configures the Zap logging system based on configuration
func SetupLogging(cfg config.LoggingConfig) (*zap.Logger, error) {
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
		return nil, err
	}

	// Set global logger
	zap.ReplaceGlobals(logger)

	return logger, nil
}

// GetLogger returns the global logger instance
func GetLogger() *zap.Logger {
	return zap.L()
}

// GetSugaredLogger returns the global sugared logger instance
func GetSugaredLogger() *zap.SugaredLogger {
	return zap.S()
}
