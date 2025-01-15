package services

import (
	//"digital-library/models"
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


func (s AdminService) DeleteBook(Title string) (string, error) {
	
	book, err := s.adminRepo.FindIdByBookName(Title)
	if err != nil {
		return "", errors.New("Book not found")
	}

	err = s.adminRepo.DeleteBook(book)
	if err != nil {
		return "Book Succesfully Deleted", errors.New("Unable to delete book")
	}
	return "Book Successfully Deleted",nil


}