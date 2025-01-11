package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(s string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash password")
	}
	return string(hashedPassword), nil
}
