package main

import (
	"digital-library/config"
	"digital-library/controllers"
	"digital-library/repositories"
	"digital-library/routes"
	"digital-library/services"

	"github.com/gin-gonic/gin"

)

func main() {
	// Initialize the database connection
	config.InitDatabase()

	// Initialize repositories
	bookRepo := repositories.NewBookRepository()
	borrowRepo := repositories.NewBorrowRepository()
	adminRepo := repositories.NewAdminRepository()
	userRepo := repositories.NewUserRepository()

	// Initialize services
	bookService := services.NewBookService(bookRepo)
	borrowService := services.NewBorrowService(borrowRepo)
	adminService := services.NewAdminService(adminRepo)
	userService := services.NewUserService(userRepo)

	// Initialize controllers
	bookController := controllers.NewBookController(bookService)
	borrowController := controllers.NewBorrowController(borrowService)
	adminController := controllers.NewAdminController(adminService)
	userController := controllers.NewUserController(userService)

	// Create a new Gin router
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r, bookController, borrowController, adminController, userController)

	// Start the server

	r.Run(":8080",)
}
