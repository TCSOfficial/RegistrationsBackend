package database

import (
	"gorm.io/gorm"
  "gorm.io/driver/mysql"
	"os"
	"log"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(os.Getenv("dsn")), &gorm.Config{})

	if err != nil {
		log.Printf("Error while connecting to database: %v\n", err.Error())
		os.Exit(1)
	}
}
