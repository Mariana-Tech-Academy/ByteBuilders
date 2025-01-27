package repositories

import (
	"digital-library/config"
	"digital-library/models"
)

type UserRepository interface {
	FindUserByUsername(username string) (models.User, error)
	AddUser(user models.User) error
	AddTokenToBlacklist(tokenString models.BlacklistedToken) error
	ListBorrowedBooks(UserID uint) ([]models.Borrow, error)
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
func (r *userRepository) ListBorrowedBooks(UserID uint) ([]models.Borrow, error) {
	var existingBook []models.Borrow

	if err := config.DB.Where("user_id = ? AND status = ?", UserID, "borrowed").Find(&existingBook).Error; err != nil {
		return []models.Borrow{}, err
	}

	return existingBook, nil
}
