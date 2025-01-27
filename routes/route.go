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
	r.GET("/getauthors", userController.GetAuthors)
	// Book routes
	r.DELETE("/deletebook/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"), bookController.DeleteBook)

	// Borrow and return routes

	r.POST("/borrowbook/:id", middleware.AuthMiddleware(), borrowController.BorrowBook)

	// Admin routes for user management
	r.POST("/addbook",middleware.AuthMiddleware(),adminController.AddBook)

	r.PUT("/updatebook", middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"), adminController.UpdateBook)

	r.GET("/user", middleware.AuthMiddleware(), userController.GetUserByUsername)

	r.POST("/addauthor", middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"), adminController.AddAuthor)

}
