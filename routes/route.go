package routes

import (
	"digital-library/controllers"
	"digital-library/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine,
	bookController *controllers.BookController,
	borrowController *controllers.BorrowController,
	adminController *controllers.AdminController,
	userController *controllers.UserController) {

	// Public authentication routes
	r.POST("/signup", userController.Register)
	r.POST("/login", userController.Login)
	r.POST("/logout", userController.Logout)
	// Book routes

	// Borrow and return routes

	// Admin routes for user management

	r.GET("/user", middleware.AuthMiddleware(), userController.GetUserByUsername)

}
