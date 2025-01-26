package services

import (
	"digital-library/models"
	"digital-library/repositories"
	"errors"
)

type BorrowService interface {
	BorrowBook(bookid uint, username string) (models.Book, error)
}

type borrowService struct {
	borrowRepo repositories.BorrowRepository
	userRepo   repositories.UserRepository
}

func NewBorrowService(repo repositories.BorrowRepository, usrRepo repositories.UserRepository) BorrowService {
	return &borrowService{
		userRepo:   usrRepo,
		borrowRepo: repo,
	}
}

func (p *borrowService) BorrowBook(bookid uint, username string) (models.Book, error) {

	user, err := p.userRepo.FindUserByUsername(username)
	if err != nil {
		return models.Book{}, err
	}

	book, err := p.borrowRepo.FindBookByBookID(bookid)
	if err != nil {
		return models.Book{}, err
	}

	if book.Copies <= 0 {
		return models.Book{}, errors.New("not enugh books to borrow")
	}

	book.Copies--

	if book.Copies <= 0 {
		book.Available = false
	} else {
		book.Available = true
	}

	_, err = p.borrowRepo.UpdateBook(book)
	if err != nil {
		return models.Book{}, err
	}

	borrowRecord := models.Borrow{
		UserID: user.ID,
		BookID: bookid,
		Status: "borrowed",
	}

	err = p.borrowRepo.CreateBorrow(borrowRecord)
	if err != nil {
		return models.Book{}, errors.New("failed to create borrow entry record")
	}
	return book, nil
}
