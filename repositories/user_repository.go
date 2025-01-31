package repositories

import (
	"digital-library/config"
	"digital-library/models"
	"errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByUsername(username string) (models.User, error)
	AddUser(user models.User) error
	AddTokenToBlacklist(tokenString models.BlacklistedToken) error
	ListBorrowedBooks(UserID uint) ([]models.Borrow, error)
	GetAllAuthors() ([]models.Author, error)
	FindUserIDByUsername(username string) (uint, error)
	CreateBorrow(borrow models.Borrow) error
	FindBookByBookID(bookid uint) (models.Book, error)
	UpdateBook(book models.Book) (models.Book, error)
	FindBookByEntry(search string) ([]models.Book, error)
	FindBorrowedRecordByBookID(bookID uint) (models.Borrow, error)
	MarkBorrowAsReturned(borrowID uint) error
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

func (r *userRepository) GetAllAuthors() ([]models.Author, error) {
	var authors []models.Author

	err := config.DB.Find(&authors).Error

	if err != nil {
		return nil, err
	}
	return authors, nil

}

func (r *userRepository) FindUserIDByUsername(username string) (uint, error) {
	var existingUser models.User

	if err := config.DB.Where("username = ?", username).First(&existingUser).Error; err != nil {
		return 0, err
	}
	return existingUser.ID, nil
}

func (r *userRepository) FindBookByBookID(bookid uint) (models.Book, error) {
	var existingBook models.Book

	err := config.DB.Where("id = ?", bookid).First(&existingBook).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.Book{}, errors.New("book not found")
		}
		return models.Book{}, err
	}
	return existingBook, nil
}

func (r *userRepository) UpdateBook(book models.Book) (models.Book, error) {
	if err := config.DB.Save(&book).Error; err != nil {
		return models.Book{}, nil
	}
	return book, nil
}

func (r *userRepository) CreateBorrow(borrow models.Borrow) error {
	return config.DB.Create(&borrow).Error
}

func (r *userRepository) FindBookByEntry(search string) ([]models.Book, error) {
	var books []models.Book

	if err := config.DB.Where("title LIKE ? OR description LIKE ? OR author_name LIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%").Find(&books).Error; err != nil {
		return []models.Book{}, err
	}
	return books, nil

}

func (r *userRepository) FindBorrowedRecordByBookID(bookid uint) (models.Borrow, error) {
	var borrow models.Borrow

	err := config.DB.Where("book_id = ?", bookid).First(&borrow).Error
	if err != nil {
		return models.Borrow{}, err
	}
	return borrow, nil
}

func (r *userRepository) MarkBorrowAsReturned(borrowid uint) error {
	return config.DB.Model(&models.Borrow{}).Where("id = ?", borrowid).Update("status", "returned").Error
}
