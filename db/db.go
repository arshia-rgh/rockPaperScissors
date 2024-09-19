package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"rockPaperScissors/db/migrations"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the db: ", err)
	}

	DB = database

}

func MigrateDatabase() {
	err := migrations.CreateUsersTable(DB)

	if err != nil {
		log.Fatal("Filed to run migrations: ", err)
	}
}
