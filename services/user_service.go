package services

import (
	"digital-library/middleware"
	"digital-library/models"
	"digital-library/repositories"
	"digital-library/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp(request models.User) error
	Login(username, password string) (string, error)
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

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	// Generate JWT token
	token, err := middleware.GenerateJWT(user.Username, user.Role)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
