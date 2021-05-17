package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
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
