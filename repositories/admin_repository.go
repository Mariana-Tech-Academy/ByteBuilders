package repositories

import (
	"digital-library/config"
	"digital-library/models"
	"errors"

	"gorm.io/gorm"
)

type AdminRepository struct {
}

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

// writing the db method that adds the book to the book table in the DB

func (a *AdminRepository) AddBook(book models.Book) error {
	return config.DB.Create(&book).Error
}

func (r *AdminRepository) AddAuthorRecord(name string) (models.Author, error) {

	var existingAuthor models.Author
	if err := config.DB.Where("name = ?", name).First(&existingAuthor).Error; err != nil {
		return models.Author{}, err
	}

	author := models.Author{Name: name}
	if err := config.DB.Create(&author).Error; err != nil {
		return models.Author{}, err
	}
	return author, nil
}

// this will take in the id and delete the Book from the database
func (r *AdminRepository) DeleteBook(id uint) error {
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

func (r *AdminRepository) Deleteuser(UserID uint) (models.User, error) {

	var deletingUser models.User

	if err := config.DB.Where("id = ?", UserID).First(&deletingUser).Error; err != nil {
		return models.User{}, err
	}

	if err := config.DB.Where("id = ?", UserID).Delete(&deletingUser).Error; err != nil {
		return models.User{}, err
	}
	return deletingUser, nil
}
