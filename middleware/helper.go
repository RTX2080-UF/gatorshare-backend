package middleware

import (
	"errors"
	"log"
	"os"
	"strconv"

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

func ConvertStrToInt(idStr string) (uint, error) {
	id, err := strconv.Atoi(idStr)
    if err != nil {
		return 0,err
    }

	if id <= 0 {
		err := errors.New("invalid value for Id")
		return 0, err
	}

	return uint(id), nil
}