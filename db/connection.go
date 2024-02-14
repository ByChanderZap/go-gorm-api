package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=alex password=password dbname=gorm port=5432"
var DB *gorm.DB

func ConnectToDb() {
	var dbError error
	DB, dbError = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if dbError != nil {
		log.Fatal(dbError)
	} else {
		log.Println("DB Connected")
	}
}
