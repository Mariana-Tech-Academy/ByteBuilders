package controllers

import (
	"digital-library/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookService services.BookService
}

func NewBookController(service services.BookService) *BookController {
	return &BookController{bookService: service}
}

func (p *BookController) DeleteBook(c *gin.Context) {
	parmamID := c.Param("id")
	id, err := strconv.Atoi(parmamID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book ID not available "})
		return
	}

	err = p.bookService.DeleteBook(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "The deletion was succesful"})

}
