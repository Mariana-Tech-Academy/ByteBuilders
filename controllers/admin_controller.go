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
