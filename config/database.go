package config

import (
	"digital-library/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Println("Error loading .env file")
	}
	dsn := os.Getenv("DATABASE_URL")
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = DB.AutoMigrate(&models.Author{}, &models.Book{}, &models.User{}, &models.Borrow{})
	if err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}

	log.Println("Database connection established")
}
