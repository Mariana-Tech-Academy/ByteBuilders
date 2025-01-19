package services

import (
	"digital-library/repositories"
)

type BookService interface {
    DeleteBook(id uint) (error)
}

type bookService struct {
    bookRepo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
    return &bookService{bookRepo: repo}
}

func (s *bookService) DeleteBook(id uint) (error) {
    // delete the book by the ID 
    err := s.bookRepo.DeleteBook(id)
    if err != nil {
        return err
    }
    return nil
}
