package main

import (
	"net/http"

	"go-microservice/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/users", controllers.GetUsers)
	router.POST("/users", controllers.SaveUser)
	router.PUT("/users", controllers.UpdateUser)
	router.DELETE("/users", controllers.DeleteUser)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
