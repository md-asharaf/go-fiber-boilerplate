package auth

import (
	"errors"
	"time"

	"github.com/md-asharaf/go-fiber-boilerplate/internal/models"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services/email"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services/jwt"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services/otp"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services/redis"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/utils"

	"gorm.io/gorm"
)

// AuthService handles authentication business logic
type AuthService struct {
	db           *gorm.DB
	jwtService   *jwt.JWTService
	redisService *redis.RedisService
	emailService *email.EmailService
	otpService   *otp.OtpService
}

// NewAuthService creates a new auth service
func NewAuthService(db *gorm.DB, jwtService *jwt.JWTService, redisService *redis.RedisService, emailService *email.EmailService, otpService *otp.OtpService) *AuthService {
	return &AuthService{
		db:           db,
		jwtService:   jwtService,
		redisService: redisService,
		emailService: emailService,
		otpService:   otpService,
	}
}

// Register creates a new user account
func (a *AuthService) Register(input models.RegisterInput) (*models.AuthResponse, error) {
	// Check if user already exists
	var existingUser models.User
	if err := a.db.Where("email = ? OR username = ?", input.Email, input.Username).First(&existingUser).Error; err == nil {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := utils.Hash(input.Password)
	if err != nil {
		return nil, err
	}

	// Create user
	user := models.User{
		Email:     input.Email,
		Username:  input.Username,
		Password:  hashedPassword,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		IsActive:  true,
	}

	if err := a.db.Create(&user).Error; err != nil {
		return nil, err
	}

	// Generate access token (24 hours)
	accessToken, err := a.jwtService.GenerateToken(&user, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	// Generate refresh token (7 days)
	refreshToken, err := a.jwtService.GenerateToken(&user, 7*24*time.Hour)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		User:         user.ToResponse(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
	}, nil
}

// Login authenticates a user
func (a *AuthService) Login(input models.LoginInput) (*models.AuthResponse, error) {
	// Find user
	var user models.User
	if err := a.db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Check if user is active
	if !user.IsActive {
		return nil, errors.New("account is inactive")
	}

	// Verify password
	valid, err := utils.Verify(input.Password, user.Password)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, errors.New("invalid credentials")
	}

	// Generate access token (24 hours)
	accessToken, err := a.jwtService.GenerateToken(&user, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	// Generate refresh token (7 days)
	refreshToken, err := a.jwtService.GenerateToken(&user, 7*24*time.Hour)
	if err != nil {
		return nil, err
	}

	return &models.AuthResponse{
		User:         user.ToResponse(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
	}, nil
}
