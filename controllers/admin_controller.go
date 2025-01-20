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
	
		msg, err := r.adminService.UpdateBook(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	
		
	
		c.JSON(http.StatusOK, gin.H{"message": msg})
	}