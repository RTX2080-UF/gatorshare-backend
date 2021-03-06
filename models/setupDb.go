package models

import (
	"fmt"
	"gatorshare/middleware"
	"log"
	"gorm.io/gorm/logger"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabaseSqlLite(dbname string) *gorm.DB{
	if dbname == "" {
		dbname = "gorm.db"
	}

	database, err := gorm.Open(sqlite.Open(dbname), &gorm.Config{Logger: logger.Default.LogMode(logger.Error),})

	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	log.Print("Connected to database successfully")
	return database
}

func ConnectDatabasePostgres(dbinfo string) *gorm.DB{
	database, err := gorm.Open(postgres.Open(dbinfo), &gorm.Config{Logger: logger.Default.LogMode(logger.Error),})

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
		println("Connecting to database", database)
		db = ConnectDatabaseSqlLite(database)
		
		if res := db.Exec("PRAGMA foreign_keys = ON", nil); res.Error != nil {
			log.Fatal("Failed to enable foreign key!")
		}
	}
	
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Post{})
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&TagPost{})
	db.AutoMigrate(&TagUser{})
	db.AutoMigrate(&Follower{})
	db.AutoMigrate(&ResetPassword{})
	db.AutoMigrate(&UserPost{})
	db.AutoMigrate(&Notification{})
	db.AutoMigrate(&FeedBack{})
}

func GetDB() *gorm.DB {
	return db
}
