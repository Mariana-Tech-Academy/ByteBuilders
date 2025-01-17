package repositories

import (
	"digital-library/models"
	"digital-library/config"
)

type AdminRepository struct {}

func NewAdminRepository() *AdminRepository {
	return &AdminRepository{}
}
//writing the db method that adds the book to the book table in the DB
func (a *AdminRepository) AddBook(book models.Book) error {
    return config.DB.Create(&book).Error
}