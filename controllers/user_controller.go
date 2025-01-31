package controllers

import (
	"digital-library/models"
	"net/http"
	"strconv"
	"strings"

	"digital-library/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{userService: service}
}

func (u *UserController) Register(c *gin.Context) {
	var req models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := u.userService.SignUp(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})
}

func (u *UserController) Login(c *gin.Context) {
	var req models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := u.userService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "message": "Login successful"})
}

func (u *UserController) GetUserByUsername(c *gin.Context) {

	username := c.Query("username")

	user, err := u.userService.GetUserByUserName(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	user.Password = ""

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (u *UserController) Logout(c *gin.Context) {

	tokenWithBearer := c.GetHeader("Authorization")
	if tokenWithBearer == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
		c.Abort()
		return
	}

	// Split token from Bearer
	parts := strings.Split(tokenWithBearer, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		c.Abort()
		return
	}

	tokenString := parts[1]

	err := u.userService.Logout(tokenString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "message": "Logout successful"})
}

func (u *UserController) ListBorrowedBooks(c *gin.Context) {

	username := c.GetString("username")

	borrowedbooks, err := u.userService.ListBorrowedBooks(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": borrowedbooks})

}

func (u *UserController) GetAuthors(ctx *gin.Context) {
	authors, err := u.userService.GetAllAuthors()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	//response

	ctx.JSON(http.StatusOK, gin.H{"message": authors})
}

func (u *UserController) BorrowBook(c *gin.Context) {

	username := c.GetString("username")

	paramsid := c.Param("id")
	id, err := strconv.Atoi(paramsid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book ID not available"})
		return
	}

	book, err := u.userService.BorrowBook(uint(id), username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": book})
}

func (t *UserController) SearchBooks(c *gin.Context) {

	query := c.Query("q")

	books, err := t.userService.SearchForBook(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to find book"})
	}

	c.JSON(http.StatusFound, gin.H{"message": books})
}

func (u *UserController) ReturnBorrowedBook(c *gin.Context) {

	paramsid := c.Param("id")
	id, err := strconv.Atoi(paramsid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book ID not available"})
		return
	}
	err = u.userService.ReturnBorrowedBook(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book returned"})
}
