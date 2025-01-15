package controllers

import (
	"digital-library/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	adminService *services.AdminService
}

func NewAdminController(adminService *services.AdminService) *AdminController {
	return &AdminController{adminService: adminService}
}

func (p *AdminController) DeleteBook(c *gin.Context) {

	Title := c.Query("Title")

	deleteMessage, err := p.adminService.DeleteBook(Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": deleteMessage})

}

