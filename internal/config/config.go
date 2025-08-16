package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	JWT      JWTConfig
	SMTP     SMTPConfig
	Logger   LoggerConfig
	App      AppConfig
}

// ServerConfig holds server-specific configuration
type ServerConfig struct {
	Host string
	Port int
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	URL string
}

// RedisConfig holds Redis configuration
type RedisConfig struct {
	Host     string
	Port     int
	Password string
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret string
}

// EmailConfig holds email configuration
type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

// LoggingConfig holds logging configuration
type LoggerConfig struct {
	Level string
}

// AppConfig holds application-specific configuration
type AppConfig struct {
	Environment string
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	// Load .env file (check for errors)
	err := godotenv.Load(".env", ".env.local")
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg := &Config{}
	var errs []error

	// Server Config
	cfg.Server.Host = getEnv("SERVER_HOST", "localhost") // This is optional, but common practice
	if cfg.Server.Port, err = getEnvAsIntRequired("SERVER_PORT"); err != nil {
		errs = append(errs, fmt.Errorf("server port: %w", err))
	}

	// Database Config - All are required
	if cfg.Database.URL, err = getEnvRequired("DATABASE_URL"); err != nil {
		errs = append(errs, fmt.Errorf("database url: %w", err))
	}

	// Redis Config
	if cfg.Redis.Host, err = getEnvRequired("REDIS_HOST"); err != nil {
		errs = append(errs, fmt.Errorf("redis host: %w", err))
	}
	if cfg.Redis.Port, err = getEnvAsIntRequired("REDIS_PORT"); err != nil {
		errs = append(errs, fmt.Errorf("redis port: %w", err))
	}
	cfg.Redis.Password = getEnv("REDIS_PASSWORD", "") // Password can be empty

	// JWT Config - All are required
	if cfg.JWT.Secret, err = getEnvRequired("JWT_SECRET"); err != nil {
		errs = append(errs, fmt.Errorf("jwt secret: %w", err))
	}

	// SMTP Config - All are required
	if cfg.SMTP.Host, err = getEnvRequired("SMTP_HOST"); err != nil {
		errs = append(errs, fmt.Errorf("smtp host: %w", err))
	}
	if cfg.SMTP.Port, err = getEnvAsIntRequired("SMTP_PORT"); err != nil {
		errs = append(errs, fmt.Errorf("smtp port: %w", err))
	}
	if cfg.SMTP.Username, err = getEnvRequired("SMTP_USERNAME"); err != nil {
		errs = append(errs, fmt.Errorf("smtp username: %w", err))
	}
	if cfg.SMTP.Password, err = getEnvRequired("SMTP_PASSWORD"); err != nil {
		errs = append(errs, fmt.Errorf("smtp password: %w", err))
	}
	if cfg.SMTP.From, err = getEnvRequired("SMTP_FROM_EMAIL"); err != nil {
		errs = append(errs, fmt.Errorf("smtp from email: %w", err))
	}

	// Logger Config
	cfg.Logger.Level = getEnv("LOG_LEVEL", "info") // Can have a default level

	// App Config
	cfg.App.Environment = getEnv("ENV", "development") // Can have a default environment

	if len(errs) > 0 {
		return nil, errors.Join(errs...)
	}

	return cfg, nil
}

// Helper functions for optional values (with a default)
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Helper functions for required values (returns an error if missing)
func getEnvRequired(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("missing required environment variable: %s", key)
	}
	return value, nil
}

func getEnvAsIntRequired(key string) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return 0, fmt.Errorf("missing required environment variable: %s", key)
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("invalid integer value for environment variable %s: %w", key, err)
	}
	return intValue, nil
}
