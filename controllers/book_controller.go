package controllers

import (
	"digital-library/services"
)

type BookController struct {
	bookService services.BookService
}

func NewBookController(service services.BookService) *BookController {
	return &BookController{bookService: service}
}
