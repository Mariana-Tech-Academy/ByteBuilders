package repositories

import ()

type BorrowRepository interface {
}

type borrowRepository struct{}

func NewBorrowRepository() BorrowRepository {
	return &borrowRepository{}
}
