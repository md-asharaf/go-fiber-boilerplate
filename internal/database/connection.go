package database

import (
	"time"

	"github.com/md-asharaf/go-fiber-boilerplate/internal/config"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Connect establishes a connection to the database and returns the *gorm.DB instance
func Connect(cfg config.DatabaseConfig) (*gorm.DB, error) {
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
	db, err := gorm.Open(postgres.Open(cfg.URL), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		return nil, err
	}

	// Get underlying sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Configure connection pool with default values
	sqlDB.SetMaxOpenConns(25)           // Default max open connections
	sqlDB.SetMaxIdleConns(10)           // Default max idle connections
	sqlDB.SetConnMaxLifetime(time.Hour) // Default connection lifetime

	return db, nil
}

// Close closes the database connection
func Close(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

type GormLogWriter struct{}

// Printf implements the logger.Writer interface
func (w *GormLogWriter) Printf(format string, args ...interface{}) {
	zap.S().Infof(format, args...)
}
