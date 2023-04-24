package services

import (
	"context"
	"fmt"
	"go-microservice/models"
	"time"
)

type UserService struct{}

type UserQueries struct {
	SelectAll string
}

var dbs DBService

// Queries
var SelectAll string = "select * from users;"
var SelectUser string = "select * from users where id = '?';"
var InsertUser string = "insert into users values(?, ?, ?)"
var UpdateUser string = "update users set name = ?, email = ? where id = ?"
var DeleteUser string = "delete from users where id = ? "

func (us *UserService) FindUsers() []models.User {
	connection := dbs.GetConnection()
	users := []models.User{}
	rows, err := connection.Query(SelectAll)
	if err != nil {
		fmt.Println("Error while get users")
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			fmt.Println("Unable to scan user rows")
		}
		users = append(users, user)
	}

	return users
}

func (us *UserService) FindUserById(ID string) *models.User {
	connection := dbs.GetConnection()
	row := connection.QueryRow(SelectUser, ID)

	if err := row.Err(); err != nil {
		fmt.Println("Error while get users", err)
	}
	user := &models.User{}
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		fmt.Println("Unable to scan user rows")
	}

	return user
}

func (us *UserService) InsertUser(user *models.User) models.User {
	connection := dbs.GetConnection()
	fmt.Println(user)
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := connection.ExecContext(ctx, InsertUser, &user.ID, &user.Name, &user.Email)
	defer cancelfunc()
	if err != nil {
		fmt.Println("Error while insert user: ", err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error while find rows affected: ", err)
	}
	fmt.Println("rows affected:", rows)
	return *user
}

func (us *UserService) UpdateUser(user models.User) models.User {
	connection := dbs.GetConnection()
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := connection.ExecContext(ctx, UpdateUser, user.Name, user.Email, user.ID)
	defer cancelfunc()
	if err != nil {
		fmt.Println("Error while insert user: ", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error while find rows affected: ", err)
	}
	fmt.Println("rows affected:", rows)
	return user
}

func (us *UserService) DeleteUser(user models.User) models.User {
	connection := dbs.GetConnection()
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	result, err := connection.ExecContext(ctx, DeleteUser, user.ID)
	defer cancelfunc()
	if err != nil {
		fmt.Println("Error while delete user: ", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error while find rows affected: ", err)
	}
	fmt.Println("rows affected:", rows)

	return user
}
