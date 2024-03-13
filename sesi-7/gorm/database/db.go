package database

import (
	"fmt"
	"hacktiv-digitalent-golang/sesi-7/gorm/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host = "localhost"
	user = "root"
	password = "rajawali02"
	dbPort = "5432"
	dbName = "learning-gorm"
	db *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database: ", err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Product{})
}