package services

import (
	"digital-library/repositories"
)

type BorrowService interface {
}

type borrowService struct {
	borrowRepo repositories.BorrowRepository
}

func NewBorrowService(repo repositories.BorrowRepository) BorrowService {
	return &borrowService{borrowRepo: repo}
}
