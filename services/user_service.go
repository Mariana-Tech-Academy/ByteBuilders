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
	GetAllAuthors() ([]models.Author, error)
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

func (u *userService) GetAllAuthors() ([]models.Author, error){
	return u.userRepo.GetAllAuthors()
}
