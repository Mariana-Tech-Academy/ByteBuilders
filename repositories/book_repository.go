package repositories

import (
	"digital-library/config"
	"digital-library/models"
	"errors"

	"gorm.io/gorm"
)

type BookRepository interface {
	DeleteBook(id uint) error
}

type bookRepository struct {
}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}

// this will take in the id and delete the Book from the database
func (r *bookRepository) DeleteBook(id uint) error {
	book := models.Book{}

	// check if the book exists in the database
	if err := config.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("book not found")
		}
		return err
	}

	if err := config.DB.Delete(&book, id).Error; err != nil {
		return err
	}
	return nil
}
