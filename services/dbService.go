package services

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"

	"go-microservice/utils"
)

type DBService struct {
	client *sql.DB
}

func (dbs *DBService) new() *sql.DB {
	cfg := mysql.Config{
		User:   utils.DB_USER,
		Passwd: utils.DB_PASSWORD,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%v:3306", utils.DB_HOST),
		DBName: utils.DB_NAME,
	}
	client, err := sql.Open("mysql", cfg.FormatDSN())

	dbs.client = client
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	dbs.client.SetConnMaxLifetime(time.Minute * 3)
	dbs.client.SetMaxOpenConns(10)
	dbs.client.SetMaxIdleConns(10)

	return dbs.client

}

func (dbs *DBService) GetClient() *sql.DB {
	return dbs.client
}

func (dbs *DBService) init() {
	dbs.client = dbs.new()
	pingErr := dbs.client.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}
