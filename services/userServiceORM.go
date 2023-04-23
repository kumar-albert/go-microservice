package services

import (
	"fmt"
	"go-microservice/models"

	"github.com/google/uuid"
)

type UserService struct{}

func (us *UserService) FindUsers() []models.User {
	dbs := &DBService{}
	dbs.init()
	users := []models.User{}
	result := dbs.client.Find(&users)
	if result.Error != nil {
		fmt.Sprintf("Error while get users")
	}
	return users
}

func (us *UserService) InsertUser(user models.User) models.User {
	dbs := &DBService{}
	dbs.init()
	user.ID = uuid.NewString()
	result := dbs.client.Create(&user)
	if result.Error != nil {
		fmt.Errorf("Error while insert user: %v", user)
	}
	return user
}

func (us *UserService) UpdateUser(user models.User) models.User {
	dbs := &DBService{}
	dbs.init()
	result := dbs.client.Save(&user)

	if result.Error != nil {
		fmt.Errorf("Error while update user: %v", user)
	}

	return user
}

func (us *UserService) DeleteUser(user models.User) models.User {
	dbs := &DBService{}
	dbs.init()
	result := dbs.client.Where("email = ?", user.Email).Delete(&user)

	if result.Error != nil {
		fmt.Errorf("Error while insert user: %v", user)
	}

	return user
}
