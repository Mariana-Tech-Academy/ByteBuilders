package services

import (
	"digital-library/middleware"
	"digital-library/models"
	"digital-library/repositories"
	"digital-library/utils"
	"errors"
)

type UserService interface {
	SignUp(request models.User) error
	Login(username, password string) (string, error)
	GetUserByUserName(username string) (models.User, error)
	Logout(tokenstring string) error
	ListBorrowedBooks(username string) ([]models.Borrow, error)
	GetAllAuthors() ([]models.Author, error)
	BorrowBook(bookid uint, username string) (models.Book, error)
	SearchForBook(search string) ([]models.Book,error)
	ReturnBorrowedBook(bookid uint) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{userRepo: repo}
}

func (s userService) SignUp(request models.User) error {
	// Check if the username is already taken
	_, err := s.userRepo.FindUserByUsername(request.Username)
	if err == nil {
		return errors.New("username already taken")
	}

	// Hash the password
	hash, err := utils.HashPassword(request.Password)
	if err != nil {
		return errors.New("could not hash pasword")
	}

	user := models.User{
		Username: request.Username,
		Password: hash,
		Role:     request.Role,
	}

	err = s.userRepo.AddUser(user)
	if err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (s userService) Login(username, password string) (string, error) {
	user, err := s.userRepo.FindUserByUsername(username)
	if err != nil {
		return "", errors.New("user not found")
	}

	err = utils.CompareHashAndPassword(user.Password, password)
	if err != nil {
		return "invalid password", err
	}

	// Generate JWT token
	token, err := middleware.GenerateJWT(user.Username, user.Role)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

func (s *userService) GetUserByUserName(username string) (models.User, error) {

	user, err := s.userRepo.FindUserByUsername(username)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (s *userService) Logout(tokenString string) error {
	blacklist := models.BlacklistedToken{
		Token: tokenString,
	}
	err := s.userRepo.AddTokenToBlacklist(blacklist)
	if err != nil {
		return errors.New("failed to add token")
	}

	return nil

}

// call db methode
// list of authors & errors
// if we need then

func (s *userService) ListBorrowedBooks(username string) ([]models.Borrow, error) {

	user, err := s.userRepo.FindUserByUsername(username)
	if err != nil {
		return []models.Borrow{}, err
	}

	listOfBooks, err := s.userRepo.ListBorrowedBooks(user.ID)
	if err != nil {
		return []models.Borrow{}, err
	}

	return listOfBooks, nil
}

func (s *userService) GetAllAuthors() ([]models.Author, error) {
	return s.userRepo.GetAllAuthors()
}

func (s *userService) BorrowBook(bookid uint, username string) (models.Book, error) {

	user, err := s.userRepo.FindUserByUsername(username)
	if err != nil {
		return models.Book{}, err
	}

	book, err := s.userRepo.FindBookByBookID(bookid)
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

	_, err = s.userRepo.UpdateBook(book)
	if err != nil {
		return models.Book{}, err
	}

	borrowRecord := models.Borrow{
		UserID: user.ID,
		BookID: bookid,
		Status: "borrowed",
	}

	err = s.userRepo.CreateBorrow(borrowRecord)
	if err != nil {
		return models.Book{}, errors.New("failed to create borrow entry record")
	}
	return book, nil
}


func (s *userService) SearchForBook(search string) ([]models.Book,error) {

	book , err := s.userRepo.FindBookByEntry(search)
	if err != nil {
		return []models.Book{} , err
	}
	return book , err
}

func (s *userService) ReturnBorrowedBook(bookid uint)  error {

	book, err := s.userRepo.FindBookByBookID(bookid)
	if err != nil {
		return err
	}

	borrow, err := s.userRepo.FindBorrowedRecordByBookID(bookid)
	if err != nil {
		return err
	}

	err = s.userRepo.MarkBorrowAsReturned(borrow.ID)
	if err != nil {
		return err
	}

	book.Copies++

	if book.Copies <= 0 {
		book.Available = false
	} else {
		book.Available = true
	}

	_, err = s.userRepo.UpdateBook(book)
	if err != nil {
		return err
	}

	return nil
}

