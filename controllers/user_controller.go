package controllers

import (
	"fmt"
	"net/http"

	"go-microservice/models"
	"go-microservice/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	user.ID = uuid.New()
	user = userService.InsertUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "user saved successfully!",
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
		"data":    user,
		"message": "user updated successfully!",
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
		"data":    user,
		"message": "user deleted successfully!",
	})
}
