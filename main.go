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
	adminRepo := repositories.NewAdminRepository(config.DB)

	// Initialize services
	bookService := services.NewBookService(bookRepo)
	borrowService := services.NewBorrowService(borrowRepo)
	adminService := services.NewAdminService(adminRepo)

	// Initialize controllers
	bookController := controllers.NewBookController(bookService)
	borrowController := controllers.NewBorrowController(borrowService)
	adminController := controllers.NewAdminController(adminService)

	// Create a new Gin router
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r, bookController, borrowController, adminController)

	// Start the server
	r.Run(":8080")
}
