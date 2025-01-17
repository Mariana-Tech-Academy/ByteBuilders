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

func (a AdminService) AddBook(book models.Book) (string, error){
	err := a.adminRepo.AddBook(book)
	if err != nil {
		return "Book adding unsuccessful", err
	}
	return "Book Successfully Added", nil
}
