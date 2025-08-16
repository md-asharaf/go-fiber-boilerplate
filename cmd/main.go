package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/md-asharaf/go-fiber-boilerplate/cmd/server"
	r "github.com/md-asharaf/go-fiber-boilerplate/internal/api/routes"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/config"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/database"
	m "github.com/md-asharaf/go-fiber-boilerplate/internal/models"
	s "github.com/md-asharaf/go-fiber-boilerplate/internal/services"
	u "github.com/md-asharaf/go-fiber-boilerplate/internal/utils"
	"go.uber.org/zap"
)

func main() {
	//init logger
	logger := u.InitLogger()
	//Load conf
	config, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load configuration", zap.Error(err))
	}
	// init db
	db, err := database.Connect(config.Database)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	// run migrations
	if err := db.AutoMigrate(&m.User{}); err != nil {
		logger.Fatal("Failed to run database migrations", zap.Error(err))
	}
	// init redis
	redisService, err := s.NewRedisService(config.Redis)
	if err != nil {
		logger.Fatal("Failed to initialize Redis service", zap.Error(err))
	}
	// init email,jwt,otp services
	emailService := s.NewEmailService(config.SMTP)
	jwtService := s.NewJWTService(config.JWT.Secret)
	otpService := s.NewOtpService(emailService)

	userService := s.NewUserService(db)
	authService := s.NewAuthService(db, jwtService, redisService, emailService, otpService)
	// create fiber app
	app := fiber.New()
	// set up routes
	r.SetupRoutes(app, &r.Services{
		AuthService:  authService,
		UserService:  userService,
		JWTService:   jwtService,
		RedisService: redisService,
		EmailService: emailService,
		OtpService:   otpService,
	})
	// start server
	server.StartServer(app, config.Server, logger)
	// graceful shutdown and close redis and db
	defer func() {
		logger.Info("Shutting down server")
		if err := database.Close(db); err != nil {
			logger.Error("Failed to close database connection", zap.Error(err))
		}
		if err := redisService.Close(); err != nil {
			logger.Error("Failed to close Redis connection", zap.Error(err))
		}
	}()
}
