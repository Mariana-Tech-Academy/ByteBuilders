package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username" binding:"required"` // Unique username, required field
	Password string `json:"password" binding:"required"`               // Password, required field
	Role     string `json:"role" binding:"required"`                   // Role ("admin" or "user"), required field
}

type Book struct {
	gorm.Model
	Title       string `json:"title" binding:"required"`       // Title of the book, required field
	Description string `json:"description"`                    // Optional description of the book
	AuthorID    uint   `json:"author_id" binding:"required"`   // Author ID, required field
	Copies      int    `json:"copies" binding:"required"`      // Number of available copies, required field
	AuthorName  string `json:"author_name" binding:"required"` // Author name, required field
	Available   bool   `json:"available"`                      // Availability status (default: true)
}

type Author struct {
	gorm.Model
	Name  string `json:"name" binding:"required"` // Name of the author, required field
	Books []Book `gorm:"foreignKey:AuthorID"`     // Relationship with books
}

type Borrow struct {
	gorm.Model
	UserID uint   `json:"user_id" binding:"required"` // ID of the user borrowing the book, required field
	BookID uint   `json:"book_id" binding:"required"` // ID of the borrowed book, required field
	Status string `json:"status"`                     // Borrow status ("borrowed" or "returned")
	User   User   `gorm:"foreignKey:UserID"`          // Relationship with user
	Book   Book   `gorm:"foreignKey:BookID"`          // Relationship with book
}
