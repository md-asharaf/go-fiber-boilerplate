package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/md-asharaf/go-fiber-boilerplate/cmd/server"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/api/routes"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/config"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/database"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services/auth"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services/email"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services/jwt"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services/otp"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services/redis"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/services/user"
	"github.com/md-asharaf/go-fiber-boilerplate/internal/utils"
	"go.uber.org/zap"
)

func main() {
	//init logger
	logger := utils.InitLogger()
	//Load conf
	config := config.Load()
	// init db
	db, err := database.Connect(config.Database)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	// run migrations
	if err := db.AutoMigrate(); err != nil {
		logger.Fatal("Failed to run database migrations", zap.Error(err))
	}
	// init redis
	redisService, err := redis.NewRedisService(config.Redis)
	if err != nil {
		logger.Fatal("Failed to initialize Redis service", zap.Error(err))
	}
	// init email,jwt,otp services
	emailService := email.NewEmailService(config.SMTP)
	jwtService := jwt.NewJWTService(config.JWT.Secret)
	otpService := otp.NewOtpService(emailService)

	userService := user.NewUserService(db)
	authService := auth.NewAuthService(db, jwtService, redisService, emailService, otpService)
	// create fiber app
	app := fiber.New()
	// set up routes
	routes.SetupRoutes(app, &routes.Services{
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
