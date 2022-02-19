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
	if dbname == "" {
		dbname = "gorm.db"
	}

	database, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{})

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
	dbtype := middleware.GetEnv("DB_TYPE", "sqlite", envSrc)

	if dbtype == "postgres" {
		user := middleware.GetEnv("PG_USER", "postgres", envSrc)
		password := middleware.GetEnv("PG_PASSWORD", "", envSrc)
		host := middleware.GetEnv("PG_HOST", "localhost", envSrc)
		port := middleware.GetEnv("PG_PORT", "5432", envSrc)
		database := middleware.GetEnv("DB_NAME", "gatorshare", envSrc)
		
		dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=prefer",
			user,
			password,
			host,
			port,
			database,
		)

		db = ConnectDatabasePostgres(dbinfo)
	} else if dbtype == "sqlite" {
		database := middleware.GetEnv("DB_NAME", "Db/gatorshare.db", envSrc)
		db = ConnectDatabaseSqlLite(database)
	}
	
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})
}

func GetDB() *gorm.DB {
	return db
}
