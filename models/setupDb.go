package models

import (
	"fmt"
	"log"
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
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

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Init() {
	user := getEnv("PG_USER", "postgres")
	password := getEnv("PG_PASSWORD", "")
	host := getEnv("PG_HOST", "localhost")
	port := getEnv("PG_PORT", "8080")
	database := getEnv("PG_DB", "gatorshare")
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
