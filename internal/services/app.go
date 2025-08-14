package services

import (
	"errors"
	"fmt"

	"github.com/md-asharaf/go-fiber-boilerplate/internal/config"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/database"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Services holds all initialized services for global access
// Add more services as needed
type Services struct {
	DB    *gorm.DB
	REDIS *RedisService
	JWT   *JWTService
	AUTH  *AuthService
	USER  *UserService
}

// InitServices initializes all services and returns an AppServices container
func InitServices(cfg *config.Config) (*Services, error) {
	// Initialize logger
	if err := utils.InitLogger(cfg.Logging); err != nil {
		return nil, fmt.Errorf("failed to initialize logging: %w", err)
	}
	logger := utils.Logger

	// Connect to database
	db, err := database.Connect(cfg.Database)
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Initialize Redis service
	redisService, err := NewRedisService(cfg.Redis)
	if err != nil {
		logger.Error("Failed to initialize Redis service", zap.Error(err))
		return nil, fmt.Errorf("failed to initialize Redis service: %w", err)
	}

	jwtService := NewJWTService(cfg.JWT.Secret)
	authService := NewAuthService(db, jwtService, redisService)
	userService := NewUserService(db)

	return &Services{
		DB:    db,
		REDIS: redisService,
		JWT:   jwtService,
		AUTH:  authService,
		USER:  userService,
	}, nil
}

func (s *Services) CloseServices() error {
	var errs []error
	if s.DB != nil {
		if err := database.Close(s.DB); err != nil {
			errs = append(errs, err)
		}
	}
	if s.REDIS != nil {
		if err := s.REDIS.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		// Combine errors into one
		return combineErrors(errs)
	}
	return nil
}

// combineErrors combines multiple errors into one
func combineErrors(errs []error) error {
	if len(errs) == 0 {
		return nil
	}
	msg := ""
	for _, err := range errs {
		msg += err.Error() + "; "
	}
	return errors.New(msg)
}
