package middleware

import (
	"errors"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func LoadEnv(path string) bool {
	if (FileExists(path)) {	
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		return false
	}
	return true
}

func GetEnv(key, fallback string, fromFile bool) string {
	if (fromFile) {
		value := os.Getenv(key)
		return value
	} else {
		if value, ok := os.LookupEnv(key); ok {
			return value
		}
	}

	return fallback
}