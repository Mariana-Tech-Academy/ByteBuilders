package repositories

import (
	"digital-library/config"
	"digital-library/models"
	"errors"

	"gorm.io/gorm"
)

type BorrowRepository interface {
	FindUserIDByUsername(username string) (uint, error)
	CreateBorrow(borrow models.Borrow) error
	FindBookByBookID(bookid uint) (models.Book, error)
	UpdateBook(book models.Book) (models.Book, error)
}

type borrowRepository struct{}

func NewBorrowRepository() BorrowRepository {
	return &borrowRepository{}
}

func (y *borrowRepository) FindUserIDByUsername(username string) (uint, error) {
	var existingUser models.User

	if err := config.DB.Where("username = ?", username).First(&existingUser).Error; err != nil {
		return 0, err
	}
	return existingUser.ID, nil
}

func (y *borrowRepository) FindBookByBookID(bookid uint) (models.Book, error) {
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

func (y *borrowRepository) UpdateBook(book models.Book) (models.Book, error) {
	if err := config.DB.Save(&book).Error; err != nil {
		return models.Book{}, nil
	}
	return book, nil
}

func (y *borrowRepository) CreateBorrow(borrow models.Borrow) error {
	return config.DB.Create(&borrow).Error
}
