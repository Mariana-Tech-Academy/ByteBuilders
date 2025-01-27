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
	adminRepo := repositories.NewAdminRepository()
	userRepo := repositories.NewUserRepository()

	// Initialize services
	adminService := services.NewAdminService(adminRepo)
	userService := services.NewUserService(userRepo)

	// Initialize controllers
	adminController := controllers.NewAdminController(adminService)
	userController := controllers.NewUserController(userService)

	// Create a new Gin router
	r := gin.Default()

	// Register routes
	routes.RegisterRoutes(r, adminController, userController)

	// Start the server

	r.Run(":8080")
}
