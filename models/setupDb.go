package models

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "Db/share-v.1.0-test.db")

	// database, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
	// 	DisableForeignKeyConstraintWhenMigrating: true,
	//   })

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	database.AutoMigrate(&User{})
	database.AutoMigrate(&Post{})
	database.AutoMigrate(&Comment{})
	log.Print("Connected to database successfully")
	DB = database
}