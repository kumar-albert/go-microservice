package services

import (
	"fmt"
	"go-microservice/models"
)

type UserService struct{}

func (us *UserService) FindUsers() []models.User {
	dbs := &DBService{}
	dbs.init()
	users := []models.User{}
	rows, err := dbs.client.Query("select * from user")
	if err != nil {
		fmt.Sprintf("Error while get users")
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Name, &user.Email); err != nil {
			fmt.Errorf("Unable to scan user rows")
		}
		users = append(users, user)
	}

	fmt.Println(dbs.client)
	return users
}

func (us *UserService) InsertUser(user models.User) models.User {
	dbs := &DBService{}
	dbs.init()
	row, err := dbs.client.Query("insert into user values(?, ?)", user.Name, user.Email)
	defer row.Close()
	if err != nil {
		fmt.Errorf("Error while insert user: %v", user)
	}
	userRow := models.User{}
	row.Scan(&userRow.Name, &userRow.Email)

	return userRow
}

func (us *UserService) UpdateUser(user models.User) models.User {
	dbs := &DBService{}
	dbs.init()
	row, err := dbs.client.Query("update user set name = '?' where email = '?'", user.Name, user.Email)
	defer row.Close()
	row.Scan(&user.Name, &user.Email)

	if err != nil {
		fmt.Errorf("Error while insert user: %v", user)
	}

	return user
}

func (us *UserService) DeleteUser(user models.User) models.User {
	dbs := &DBService{}
	dbs.init()
	row, err := dbs.client.Query("delete from user where email = '?'", user.Email)
	defer row.Close()
	row.Scan(&user.Name, &user.Email)

	if err != nil {
		fmt.Errorf("Error while insert user: %v", user)
	}

	return user
}
