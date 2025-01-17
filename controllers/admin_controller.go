package controllers

import (
	"digital-library/models"
	"digital-library/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	adminService *services.AdminService
}

func NewAdminController(service *services.AdminService) *AdminController {
	return &AdminController{adminService: service}
}

func (r *AdminController) UpdateBook(c *gin.Context) {

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

	updatedBook, err := r.adminService.UpdateBook(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": updatedBook})
}

func (a *AdminController) AddBook(ctx *gin.Context) {
	//create a request body
	var req models.Book

	//decode request body into a struct
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	//call the service layer(AddBook)
    msg ,err := a.adminService.AddBook(req)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
	//response
	
	ctx.JSON(http.StatusOK, gin.H{"message": msg})
}
