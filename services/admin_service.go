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

func (r *AdminService) UpdateBook(book models.Book) (string, error) {

	_, err := r.adminRepo.UpdateBook(book)
	if err != nil {
		return "Book have not updated", err 
	}

	return "Book updated successfuly", nil
}
