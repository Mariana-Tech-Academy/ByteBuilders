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

func (s *AdminService) UpdateBook(book models.Book) (models.Book, error) {
	author, err := s.adminRepo.AuthorExists(book.AuthorName)
	if err != nil {
		return models.Book{}, err
	}
	book.AuthorID = author.ID

	updatedBook, err := s.adminRepo.UpdateBook(book)
	if err != nil {
		return models.Book{}, err
	}

	return updatedBook, nil
}

func (s AdminService) AddBook(book models.Book) error {
	author, err := s.adminRepo.AuthorExists(book.AuthorName)
	if err != nil {
		return err
	}
	book.AuthorID = author.ID

	err = s.adminRepo.AddBook(book)
	if err != nil {
		return err
	}
	return nil
}

func (s AdminService) AddAuthor(request models.Author) (models.Author, error) {

	author, err := s.adminRepo.AddAuthorRecord(request.Name)
	if err != nil {
		return models.Author{}, err
	}
	return author, nil
}

func (s *AdminService) DeleteBook(id uint) error {
	// delete the book by the ID
	err := s.adminRepo.DeleteBook(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *AdminService) DeleteUser(UserID uint) (models.User, error) {

	user, err := s.adminRepo.Deleteuser(UserID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil

}
