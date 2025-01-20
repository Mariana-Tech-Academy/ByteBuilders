package repositories

import (
	"digital-library/config"
	"digital-library/models"
	"errors"

	"gorm.io/gorm"
)

type AdminRepository struct{}

func NewAdminRepository() *AdminRepository {
	return &AdminRepository{}
}

func (r *AdminRepository) UpdateBook(book models.Book) (models.Book, error) {

	if err := config.DB.Save(&book).Error; err != nil {
		return models.Book{}, err
	}
	return book, nil
}

func (r *AdminRepository) AuthorExists(AuthorName string) (models.Author, error) {
	// Check if the author exists
	var author models.Author
	err := config.DB.Where("name = ?", AuthorName).First(&author).Error
	if err != nil {
		// If the error is not a "record not found" error, return it
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return author, err
		}

		// Create the author if they don't exist
		author = models.Author{Name: AuthorName}
		if createErr := config.DB.Create(&author).Error; createErr != nil {
			return author, createErr
		}
	}

	return author, nil
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

//writing the db method that adds the book to the book table in the DB
func (a *AdminRepository) AddBook(book models.Book) error {
	return config.DB.Create(&book).Error
}

