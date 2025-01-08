package main

import (
	"log"
	"yusrilaprial/backend-api/controllers"
	"yusrilaprial/backend-api/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	router := gin.Default()

	db.ConnectDatabase()

	router.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/posts", controllers.FindPosts)

	router.POST("/posts", controllers.StorePost)

	router.GET("/posts/:id", controllers.FindPostById)

	router.PUT("/posts/:id", controllers.UpdatePost)

	router.DELETE("/posts/:id", controllers.DeletePost)

	router.Run(":8080")
}
