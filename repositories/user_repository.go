package repositories

import (
	"digital-library/config"
	"digital-library/models"
)

type UserRepository interface {
	FindUserByUsername(username string) (models.User, error)
	AddUser(user models.User) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindUserByUsername(username string) (models.User, error) {
	var existingUser models.User

	if err := config.DB.Where("username = ?", username).First(&existingUser).Error; err != nil {
		return models.User{}, err
	}
	return existingUser, nil
}

func (r *userRepository) AddUser(user models.User) error {
	return config.DB.Create(&user).Error
}
