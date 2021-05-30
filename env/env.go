package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

// Get get the environment key at the .env file.
func Get(key string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// GetInt get the environment key as int at the .env file.
func GetInt(key string) int {
	value := os.Getenv(key)
	if value != "" {
		i, err := strconv.Atoi(value)
		if err != nil {
			return 0
		}
		return i
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	i, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return i
}
