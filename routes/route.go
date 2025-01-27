package routes

import (
	"digital-library/controllers"
	"digital-library/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine,
	adminController *controllers.AdminController,
	userController *controllers.UserController) {

	// Public authentication routes
	r.POST("/signup", userController.Register)
	r.POST("/login", userController.Login)
	r.POST("/logout", userController.Logout)
	// Book routes
	r.DELETE("/deletebook/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"), bookController.DeleteBook)
	r.GET("/listborrowedbook", middleware.AuthMiddleware(), userController.ListBorrowedBooks)
	// Borrow and return routes

	// Admin routes for user management
	r.PUT("/updatebook", middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"), adminController.UpdateBook)


	// User routes
	r.GET("/user", middleware.AuthMiddleware(), userController.GetUserByUsername)
	r.POST("/borrowbook/:id", middleware.AuthMiddleware(), userController.BorrowBook)
	r.POST("/logout", middleware.AuthMiddleware(), userController.Logout)
	r.GET("/getauthors", middleware.AuthMiddleware(), userController.GetAuthors)

	// Admin routes for user management
	r.POST("/addbook", middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"), adminController.AddBook)
	r.PUT("/updatebook", middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"), adminController.UpdateBook)
	r.DELETE("/deletebook/:id", middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"), adminController.DeleteBook)
	r.POST("/addauthor", middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"), adminController.AddAuthor)

}
