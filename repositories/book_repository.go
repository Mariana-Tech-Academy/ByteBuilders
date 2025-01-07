package repositories

import ()

type BookRepository interface {
}

type bookRepository struct{}

func NewBookRepository() BookRepository {
	return &bookRepository{}
}
