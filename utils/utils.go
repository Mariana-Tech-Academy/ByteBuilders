package utils

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(s string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash password")
	}
	return string(hashedPassword), nil
}

func CompareHashAndPassword(existingPassword string, loginPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(existingPassword), []byte(loginPassword))
	if err != nil {
		return errors.New("invalid passowrd/ password mismatch")
	}
	return nil
}

func GetTodaysDate() string {
	today := time.Now()
	return today.Format("2006-01-02")
}

func ValidatePassword(password string) bool {

	if len(password) < 0 {
		return false
	}
	if len(password) > 8 {
		return false
	}

	return true
}
