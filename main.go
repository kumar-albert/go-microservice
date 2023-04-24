package main

import (
	"net/http"

	"go-microservice/controllers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	v1 := router.Group("/v1")
	{
		v1.GET("/users", controllers.GetUsers)
		v1.POST("/users", controllers.SaveUser)
		v1.PUT("/users", controllers.UpdateUser)
		v1.DELETE("/users", controllers.DeleteUser)
	}

	return router
}

func main() {
	r := setupRouter()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
