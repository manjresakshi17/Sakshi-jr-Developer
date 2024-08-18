// main.go
package main

import (
	"github.com/gin-gonic/gin"
	//"gorm.io/driver/postgres"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"your_project/config"
	"your_project/controllers"
	"your_project/middleware"
	"your_project/models"
)

func main() {
	r := gin.Default()

	// Initialize configuration
	config.InitConfig()

	// Connect to the database
	dsn := "host=localhost user=postgres password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	config.DB = db

	// Migrate the schema
	db.AutoMigrate(&models.User{}, &models.Post{})

	// Define routes
	r.POST("/auth/signup", controllers.Signup)
	r.POST("/auth/login", controllers.Login)

	protected := r.Group("/")
	protected.Use(middleware.Auth())
	protected.POST("/auth/logout", controllers.Logout)
	protected.POST("/posts", controllers.CreatePost)
	protected.PUT("/posts/:id", controllers.UpdatePost)
	protected.DELETE("/posts/:id", controllers.DeletePost)

	// Start the server
	r.Run(":8080")
}
