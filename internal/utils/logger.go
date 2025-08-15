package utils

import (
	"os"

	"go.uber.org/zap"
)

// InitLogger initializes the global Logger variable
func InitLogger() *zap.Logger {
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
	level, err := zap.ParseAtomicLevel(os.Getenv("LOG_LEVEL"))
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
		return zap.L()
	}
	return logger
}
