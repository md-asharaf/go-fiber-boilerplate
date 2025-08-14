package services

import (
	"github.com/google/uuid"
	"github.com/yourusername/go-backend-boilerplate/internal/models"
	"gorm.io/gorm"
)

// UserService handles user management operations
type UserService struct {
	db *gorm.DB
}

// NewUserService creates a new user service
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

// GetUserByID retrieves a user by ID
func (u *UserService) GetUserByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	if err := u.db.Where("id = ? AND is_active = ?", userID, true).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates user information
func (u *UserService) UpdateUser(userID uuid.UUID, firstName, lastName string) (*models.User, error) {
	var user models.User
	if err := u.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}

	user.FirstName = firstName
	user.LastName = lastName

	if err := u.db.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser soft deletes a user
func (u *UserService) DeleteUser(userID uuid.UUID) error {
	return u.db.Where("id = ?", userID).Delete(&models.User{}).Error
}

// ListUsers retrieves all active users with pagination
func (u *UserService) ListUsers(limit, offset int) ([]models.User, error) {
	var users []models.User
	err := u.db.Where("is_active = ?", true).
		Limit(limit).
		Offset(offset).
		Find(&users).Error
	return users, err
}
