package database

import (
	"fmt"
	"hacktiv-digitalent-golang/sesi-10/securing-app-with-jwt/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "rajawali02"
	dbPort   = "5432"
	dbname   = "apiv2"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("error connecting to database:", err)
		return // exit function if error occurs
	}

	fmt.Println("Database connected successfully")
	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
