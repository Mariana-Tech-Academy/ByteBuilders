package repositories

import (
	"digital-library/config"
	"digital-library/models"
)

type AdminRepository struct {}

func NewAdminRepository() *AdminRepository {
	return &AdminRepository{}
}

// writing the db method that adds the book to the book table in the DB
func (a *AdminRepository) AddBook(book models.Book) error {
	return config.DB.Create(&book).Error
}
