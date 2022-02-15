package models

import (
	"fmt"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"gatorshare/middleware"
)

var db *gorm.DB

func ConnectDatabaseSqlLite(dbname string) *gorm.DB{
	database, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	log.Print("Connected to database successfully")
	return database
}

func ConnectDatabasePostgres(dbinfo string) *gorm.DB{
	database, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	log.Print("Connected to database successfully")
	return database
}

func Init(envSrc bool) {
	user := middleware.GetEnv("PG_USER", "postgres", envSrc)
	password := middleware.GetEnv("PG_PASSWORD", "", envSrc)
	host := middleware.GetEnv("PG_HOST", "localhost", envSrc)
	port := middleware.GetEnv("PG_PORT", "5432", envSrc)
	database := middleware.GetEnv("PG_DB", "gatorshare", envSrc)
	// databaseSqlLite :=  getEnv("SQLite_DB", "Db/share-v.1.0-test.db")

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	db = ConnectDatabasePostgres(dbinfo)

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})
}

func GetDB() *gorm.DB {
	return db
}
