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

func (a AdminService) AddAuthor(request models.Author) (models.Author,error) {

	author,err := a.adminRepo.AddAuthorRecord(request.Name)
	if err != nil {
		return models.Author{},err
	}
	return author , nil
}
