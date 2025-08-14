package database

import (
	"time"

	"github.com/md-asharaf/go-fiber-boilerplate/internal/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB holds the database connection
var DB *gorm.DB

// Connect establishes a connection to the database
func Connect(cfg config.DatabaseConfig) error {
	var err error

	// Configure GORM logger
	gormLogger := logger.New(
		&GormLogWriter{},
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	// Connect to database
	DB, err = gorm.Open(postgres.Open(cfg.URL), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return err
	}

	// Get underlying sql.DB
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	// Configure connection pool with default values
	sqlDB.SetMaxOpenConns(25)           // Default max open connections
	sqlDB.SetMaxIdleConns(10)           // Default max idle connections
	sqlDB.SetConnMaxLifetime(time.Hour) // Default connection lifetime

	zap.L().Info("Database connection established")
	return nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}

// Close closes the database connection
func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// GormLogWriter implements logger.Writer interface for Zap
type GormLogWriter struct{}

// Printf implements the logger.Writer interface
func (w *GormLogWriter) Printf(format string, args ...interface{}) {
	zap.S().Infof(format, args...)
}
