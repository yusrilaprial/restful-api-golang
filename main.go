package main

import (
	"yusrilaprial/backend-api/controllers"
	"yusrilaprial/backend-api/db"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db.ConnectDataBase()

	router.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/posts", controllers.FindPosts)

	router.POST("/posts", controllers.StorePost)

	router.Run(":3000")
}
