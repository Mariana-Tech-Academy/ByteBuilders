package controllers

import (
	"digital-library/services"
)

type BorrowController struct {
	borrowService services.BorrowService
}

func NewBorrowController(service services.BorrowService) *BorrowController {
	return &BorrowController{borrowService: service}
}
