package services

import (
	"digital-library/models"
	"digital-library/repositories"
	"errors"
)

type BorrowService interface {
	BorrowBook(bookid uint, username string) (string, error)
}

type borrowService struct {
	borrowRepo repositories.BorrowRepository
}

func NewBorrowService(repo repositories.BorrowRepository) BorrowService {
	return &borrowService{borrowRepo: repo}
}

func (p *borrowService) BorrowBook(bookid uint, username string) (string, error) {

	_, err := p.borrowRepo.CheckIfBookExists(bookid)
	if err != nil {
		return "Book does not exist", err
	}

	userId, err := p.borrowRepo.FindUserIDByUsername(username)
	if err != nil {
		return "User does not exist", err
	}

	book, err := p.borrowRepo.FindBookByBookID(bookid)
	if err != nil {
		return "Book record does not exist", err
	}

	if book.Copies <= 0 {
		return "Not enough books to borrow", errors.New("not enugh books to borrow")
	}

	book.Copies -= 1

	if book.Copies <= 0 {
		book.Available = false
	} else {
		book.Available = true
	}

	_, err = p.borrowRepo.UpdateBook(book)
	if err != nil {
		return "", err
	}

	borrowRecord := models.Borrow{
		UserID: userId,
		BookID: bookid,
		Status: "borrowed",
	}

	_, err = p.borrowRepo.CreateBorrow(borrowRecord)
	if err != nil {
		return "Failed to create record", errors.New("failed to create borrow entry record")
	}
	return "Borrow record created", nil
}

// find the user if they exist and if the author exist
