package database

import (
	"Test/Test-Crud/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() *gorm.DB {

	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {

	db_URL := "host= localhost user= postgres password= 12345 dbname= mycompany port= 5432 sslmode= disable"

	db, err := gorm.Open(postgres.Open(db_URL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nConnected to DATABASE: ", db.Name())
	db.AutoMigrate(&models.User{})

	return db
}
