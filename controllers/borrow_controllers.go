package controllers

import (
	"digital-library/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BorrowController struct {
	borrowService services.BorrowService
}

func NewBorrowController(service services.BorrowService) *BorrowController {
	return &BorrowController{borrowService: service}
}

func (p *BorrowController) BorrowBook(c *gin.Context) {

	username := c.GetString("username")

	paramsid := c.Param("id")
	id, err := strconv.Atoi(paramsid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book ID not available"})
		return 
	}



	reply, err := p.borrowService.BorrowBook(uint(id), username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": reply})
}
