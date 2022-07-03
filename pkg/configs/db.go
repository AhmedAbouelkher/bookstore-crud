package configs

import (
	"bookstore-crud/pkg/models"
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
	
	syncTables()
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

func syncTables() {
	db.AutoMigrate(&models.Book{}, &models.Author{})

	// migrations.RunMigrations()
	
	log.Println("Tables have been auto migrated")
}