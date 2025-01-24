package repositories

import (
	"digital-library/config"
	"digital-library/models"
	"errors"

	"gorm.io/gorm"
)

type BorrowRepository interface {
	CheckIfBookExists(id uint) (string, error)
	FindUserIDByUsername(username string) (uint, error)
	CreateBorrow(borrow models.Borrow) (string, error)
	FindBookByBookID(bookid uint) (models.Book, error)
	UpdateBook(book models.Book) (models.Book, error)
}

type borrowRepository struct{}

func NewBorrowRepository() BorrowRepository {
	return &borrowRepository{}
}

func (y *borrowRepository) CheckIfBookExists(id uint) (string, error) {
	var book models.Book
	err := config.DB.Where("id = ?", id).First(&book).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("book not found")
		}
		return "", err
	}
	return "", nil
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

	if err := config.DB.Where("id = ?", bookid).First(&existingBook).Error; err != nil {
		return models.Book{}, nil
	}
	return existingBook, nil
}

func (y *borrowRepository) UpdateBook(book models.Book) (models.Book, error) {
	if err := config.DB.Save(&book).Error; err != nil {
		return models.Book{}, nil
	}
	return book, nil
}

func (y *borrowRepository) CreateBorrow(borrow models.Borrow) (string, error) {
	return "", config.DB.Create(&borrow).Error
}
