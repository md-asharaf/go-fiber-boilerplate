package services

import (
	"github.com/yourusername/go-backend-boilerplate/internal/config"
	"gorm.io/gorm"
)

// AppServices holds all initialized services for global access
// Add more services as needed
type AppServices struct {
	DB                *gorm.DB
	EncryptionService *EncryptionService
	JWTService        *JWTService
	AuthService       *AuthService
	UserService       *UserService
}

// InitServices initializes all services and returns an AppServices container
func InitServices(cfg *config.Config, db *gorm.DB) *AppServices {
	encryptionService := NewEncryptionService(cfg.JWT.Secret)
	jwtService := NewJWTService(cfg.JWT.Secret)
	authService := NewAuthService(db, encryptionService, jwtService)
	userService := NewUserService(db)

	return &AppServices{
		DB:                db,
		EncryptionService: encryptionService,
		JWTService:        jwtService,
		AuthService:       authService,
		UserService:       userService,
	}
}
