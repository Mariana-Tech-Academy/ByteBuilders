package repositories

import (
	"digital-library/config"
	"digital-library/models"
)

type AdminRepository struct{}

func NewAdminRepository() *AdminRepository {
	return &AdminRepository{}
}

func (r *AdminRepository) AuthorExists(AuthorName string) (models.Author, error) {
	// Check if the author exists
	var author models.Author
	if err := config.DB.Where("name = ?", AuthorName).First(&author).Error; err != nil {
		// Create the author if they don't exist
		author = models.Author{Name: AuthorName}
		config.DB.Create(&author)
	}
	return author, nil
}

// writing the db method that adds the book to the book table in the DB
func (a *AdminRepository) AddBook(book models.Book) error {
	return config.DB.Create(&book).Error
}
