package repositories

import (
	"digital-library/config"
	"digital-library/models"
)

type UserRepository interface {
	FindUserByUsername(username string) (models.User, error)
	AddUser(user models.User) error
	AddTokenToBlacklist(tokenString models.BlacklistedToken) error
	GetAllAuthors() ([]models.Author, error)
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

func (r *userRepository) AddTokenToBlacklist(tokenString models.BlacklistedToken) error {
	return config.DB.Create(&tokenString).Error

}

func (r *userRepository) GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author

	err := config.DB.Find(&authors).Error

	if err != nil {
		return nil, err
	}
	return authors, nil

}
