package configs

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectToDatabase() {
	database, err := gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})

	if err != nil  {
		panic(err)
	}
	
	log.Println("Database is up and running...")
	db = database
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() error {
	sql, err := db.DB()
	if err != nil {
		return err
	}

	return sql.Close()
}