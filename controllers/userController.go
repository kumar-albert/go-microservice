package controllers

import (
	"fmt"
	"net/http"

	"go-microservice/models"
	"go-microservice/services"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	userService := &services.UserService{}
	users := userService.FindUsers()
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func SaveUser(c *gin.Context) {
	var user models.User
	userService := &services.UserService{}
	if err := c.BindJSON(&user); err != nil {
		fmt.Errorf("Error while bind user request")
	}
	user = userService.InsertUser(user)
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	userService := &services.UserService{}
	if err := c.BindJSON(&user); err != nil {
		fmt.Errorf("Error while bind user request")
	}
	user = userService.UpdateUser(user)

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	userService := &services.UserService{}
	if err := c.BindJSON(&user); err != nil {
		fmt.Errorf("Error while bind user request")
	}
	user = userService.DeleteUser(user)
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
