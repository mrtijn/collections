package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type dbConfig struct {
	db      *gorm.DB
	dbError error
}

// Config database config

// Connect to the database
func Connect() *gorm.DB {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Could not load dot env file")
	}

	db, err := gorm.Open("mysql", os.Getenv("dbConnection")+`?parseTime=true`)

	if err != nil {
		panic("failed to connect to db" + fmt.Sprintln(err))
	}

	db.LogMode(true)
	// Config.db.AutoMigrate(&models.User{})

	return db
}
