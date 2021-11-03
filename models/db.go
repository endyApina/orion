package models

import (
	"errors"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres driver
	"github.com/joho/godotenv"
)

var conn *gorm.DB

func init() {
	fmt.Println("connecting to database...")
	err := godotenv.Load("conf.env")
	if err != nil {
		LogError(errors.New("error accessing config file"))
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")

	if os.Getenv("app_mode") == "prod" {
		username = os.Getenv("prod_db_user")
		password = os.Getenv("prod_db_pass")
	}

	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbURL := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
		defer recover()
		LogError(err)
		panic(err)
	}

	conn = db
	if err == nil {
		fmt.Println("Database connection successful")
	}
	autoMigrateTables()
}

func autoMigrateTables() {
	conn.AutoMigrate(&Subscribe{})
}

//GetDB sends the db objects
func GetDB() *gorm.DB {
	return conn
}
