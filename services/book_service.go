package services

import (
	"digital-library/repositories"
)

type BookService interface {
}

type bookService struct {
	bookRepo repositories.BookRepository
}

func NewBookService(repo repositories.BookRepository) BookService {
	return &bookService{bookRepo: repo}
}
