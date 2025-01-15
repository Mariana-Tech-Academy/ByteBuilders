package repositories

import (
	"digital-library/config"
	"digital-library/models"

	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}


// this function will find the book title in the database which comes with the ID and return it
func (r *AdminRepository) FindIdByBookName(Title string) (models.Book, error) {
	var existingBook models.Book

	if err := config.DB.Where("Title = ?",Title).First(&existingBook).Error; err != nil {
		return models.Book{} , err
	}
	return existingBook, nil
}

// this will take in the id and delete the Book from the database
func (r *AdminRepository) DeleteBook(book models.Book) error {
	return config.DB.Delete(&book, book.ID).Error
}

