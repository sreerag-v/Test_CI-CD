package database

import (
	"Test/Test-Crud/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() *gorm.DB {

	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db_URL := os.Getenv("DNS")

	db, err := gorm.Open(postgres.Open(db_URL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nConnected to DATABASE: ", db.Name())
	db.AutoMigrate(&models.User{})

	return db
}
