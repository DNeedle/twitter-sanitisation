package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Pg *gorm.DB
var Connected bool = false

var dsn string = getDbUrl()

func ConnectToDatabase() {
	pg, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Pg = pg
	Connected = true

}
func GetDBConnection() *gorm.DB {
	if !Connected {
		ConnectToDatabase()
	}
	return Pg
}

func getDbUrl() string {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v", host, port, user, password, dbName)
}

func IsRecordNotFound(err error) bool {
	return err.Error() == "record not found"
}

func IsForeignKeyNotFound(err error) bool {
	return strings.Contains(err.Error(), "(SQLSTATE 23503)")

}

func Migrate() {
	if !Connected {
		ConnectToDatabase()
	}

	err := Pg.AutoMigrate(&User{})

	if err != nil {
		log.Fatal(err)
	}

}
