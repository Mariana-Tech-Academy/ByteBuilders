package services

import (
	"digital-library/models"
	"digital-library/repositories"
)

type AdminService struct {
	adminRepo *repositories.AdminRepository
}

func NewAdminService(adminRepo *repositories.AdminRepository) *AdminService {
	return &AdminService{adminRepo: adminRepo}
}

func (a AdminService) AddBook(book models.Book) error {
	author, err := a.adminRepo.AuthorExists(book.AuthorName)
	if err != nil {
		return err
	}
	book.AuthorID = author.ID

	err = a.adminRepo.AddBook(book)
	if err != nil {
		return err
	}
	return nil
}
