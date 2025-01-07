package routes

import (
	"digital-library/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine,
	bookController *controllers.BookController,
	borrowController *controllers.BorrowController,
	adminController *controllers.AdminController) {

	// Public authentication routes
	r.POST("/signup", controllers.Register)
	r.POST("/login", controllers.Login)

	// Book routes

	// Borrow and return routes

	// Admin routes for user management

}
