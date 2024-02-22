package todo

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer() {
	// Initialize Gin router
	router := gin.Default()

	// Initialize GORM
	var err error
	dsn := "host=localhost user=sayak password=pass1234 dbname=tododb port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&Todo{})
	if err != nil {
		panic("failed to migrate database schema")
	}

	// Create
	router.POST("/todos", createTodo)

	// Read all
	router.GET("/todos", getTodos)

	// Update
	router.PUT("/todos/:id", updateTodo)

	// Delete
	router.DELETE("/todos/:id", deleteTodo)

	// Run server
	router.Run(":8080")
}
