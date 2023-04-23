package services

import (
	"fmt"

	// "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-microservice/models"
	"go-microservice/utils"
)

type DBService struct {
	client *gorm.DB
}

func (dbs *DBService) new() *gorm.DB {
	gormConfig := &gorm.Config{}
	dsn := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local", utils.DB_USER, utils.DB_PASSWORD, utils.DB_HOST, utils.DB_NAME)
	client, err := gorm.Open(mysql.Open(dsn), gormConfig)
	dbs.client = client
	if err != nil {
		panic(err)
	}

	return dbs.client
}

func (dbs *DBService) GetClient() *gorm.DB {
	return dbs.client
}

func (dbs *DBService) init() {
	dbs.client = dbs.new()
	dbs.client.AutoMigrate(&models.User{})
}
