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


func (r *AdminRepository) UpdateBook(book models.Book) (models.Book, error) {


	if err := config.DB.Save(&book).Error; err != nil {
		return models.Book{}, err
	}
	return book, nil
}
