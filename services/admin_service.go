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

func (a AdminService) AddBook(book models.Book) error {
	err := a.adminRepo.AddBook(book)
	if err != nil {
		return err
	}
	return nil
}
