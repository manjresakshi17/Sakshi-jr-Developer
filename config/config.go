package config

import (
	"fmt"
	"os"

	"blog-api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DB        *gorm.DB
	JWTSecret string
)

// InitConfig initializes the database connection and configuration settings
func InitConfig() {
	var err error

	// Database setup
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=sakshi98 dbname=blogapi sslmode=disable" // Default connection string
	}
	DB, err = gorm.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		os.Exit(1)
	}

	// Perform migrations
	DB.AutoMigrate(&models.User{}, &models.Post{})

	// JWT Secret setup
	JWTSecret = os.Getenv("JWT_SECRET")
	if JWTSecret == "" {
		JWTSecret = "your-default-secret" // Provide a default secret for development
	}
}
