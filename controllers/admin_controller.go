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

func NewAdminController(adminService *services.AdminService) *AdminController {
	return &AdminController{adminService: adminService}
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
