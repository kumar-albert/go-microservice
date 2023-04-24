package services

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
)

type DBConfig struct {
	Database struct {
		Host     string
		Name     string
		User     string
		Password string
	}
}

type DBService struct {
	client *sql.DB
	config DBConfig
}

func (dbs *DBService) ReadConfig() {
	// read this from contant file
	var configFile string = os.Getenv("CONFIG_FILE")

	if configFile == "" {
		configFile = "etc/configuration/local.yaml"
	}
	yfile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	unMarshalErr := yaml.Unmarshal(yfile, &dbs.config)

	if unMarshalErr != nil {
		log.Fatal(unMarshalErr)
	}
}

func (dbs *DBService) new() *sql.DB {
	dbs.ReadConfig()
	cfg := mysql.Config{
		User:   dbs.config.Database.User,
		Passwd: dbs.config.Database.Password,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%v:3306", dbs.config.Database.Host),
		DBName: dbs.config.Database.Name,
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

func (dbs *DBService) init() {
	dbs.client = dbs.new()
	pingErr := dbs.client.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func (dbs *DBService) GetConnection() *sql.DB {
	if dbs.client == nil {
		dbs.init()
	}
	return dbs.client
}
