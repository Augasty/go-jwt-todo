package todo

import (
	"github.com/Augasty/go-jwt-todo/pkg/auth"
	"github.com/Augasty/go-jwt-todo/pkg/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func StartServer() {

	// Initialize GORM
	var err error
	dsn := "host=localhost user=sayak password=pass1234 dbname=tododb port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = user.ConnectAndMigrate(db)
	if err != nil {
		panic("failed to migrate users database schema")
	}

	// Migrate the todo schema
	err = db.AutoMigrate(&Todo{})
	if err != nil {
		panic("failed to migrate todo database schema")
	}

	router := initRouter()
	// Run server
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/token", func(c *gin.Context) {
		user.GenerateToken(db, c)
	})
	router.POST("/user/register", func(c *gin.Context) {
		user.RegisterUser(db, c)
	})
	secured := router.Group("/secured").Use(auth.Auth())
	{
		secured.POST("/todos", createTodo)
		secured.GET("/todos", getTodos)
		secured.PUT("/todos/:id", updateTodo)
		secured.DELETE("/todos/:id", deleteTodo)
	}

	return router
}
