package controllers

import (
	"digital-library/models"
	"digital-library/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	adminService *services.AdminService
}

func NewAdminController(service *services.AdminService) *AdminController {
	return &AdminController{adminService: service}
}

func (a *AdminController) AddAuthor(c *gin.Context) {

	var AuthorName struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&AuthorName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	author := models.Author{
		Name: AuthorName.Name,
	}

	author, err := a.adminService.AddAuthor(author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add author"})
	}
	c.JSON(http.StatusOK, gin.H{"message": author})
}

func (a *AdminController) UpdateBook(c *gin.Context) {

	var req models.Book

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if req.Copies < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Copies must be greater than 0"})
		return
	} else {
		req.Available = true
	}

	updatedBook, err := a.adminService.UpdateBook(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": updatedBook})
}

func (a *AdminController) AddBook(ctx *gin.Context) {
	//create a request body
	var bookPayload struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description"`
		Copies      int    `json:"copies" binding:"required"`
		AuthorName  string `json:"author_name" binding:"required"`
	}

	//decode request body into a struct
	if err := ctx.ShouldBindJSON(&bookPayload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{
		Title:       bookPayload.Title,
		Description: bookPayload.Description,
		Copies:      bookPayload.Copies,
	}

	if book.Copies < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Copies must be greater than 0"})
		return
	} else {
		book.Available = true
	}

	//call the service layer(AddBook)
	err := a.adminService.AddBook(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//response

	ctx.JSON(http.StatusOK, gin.H{"message": "Book added successfully"})
}

func (a *AdminController) DeleteBook(c *gin.Context) {
	parmamID := c.Param("id")
	id, err := strconv.Atoi(parmamID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book ID not available "})
		return
	}

	err = a.adminService.DeleteBook(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "The deletion was succesful"})

}
