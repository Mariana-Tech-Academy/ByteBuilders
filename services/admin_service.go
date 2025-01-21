package services

import (
	"digital-library/models"
	"digital-library/repositories"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

type AdminService struct {
	adminRepo *repositories.AdminRepository
}

func NewAdminService(adminRepo *repositories.AdminRepository) *AdminService {
	return &AdminService{adminRepo: adminRepo}
}

func (r *AdminService) UpdateBook(book models.Book) (models.Book, error) {
	author, err := r.adminRepo.AuthorExists(book.AuthorName)
	if err != nil {
		return models.Book{}, err
	}
	book.AuthorID = author.ID

	updatedBook, err := r.adminRepo.UpdateBook(book)
	if err != nil {
		return models.Book{}, err
	}

	return updatedBook, nil
}
