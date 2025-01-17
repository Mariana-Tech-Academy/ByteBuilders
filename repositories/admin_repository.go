package repositories

import (
	"digital-library/models"
	"digital-library/config"
	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}
//writing the db method that adds the book to the book table in the DB
func (a *AdminRepository) AddBook(book models.Book) error {
    return config.DB.Create(&book).Error
}