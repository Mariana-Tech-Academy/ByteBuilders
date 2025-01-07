package services

import (
	"digital-library/repositories"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

type AdminService struct {
	adminRepo *repositories.AdminRepository
}

func NewAdminService(adminRepo *repositories.AdminRepository) *AdminService {
	return &AdminService{adminRepo: adminRepo}
}
