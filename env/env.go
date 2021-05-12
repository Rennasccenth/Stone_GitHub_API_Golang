package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Get the key after reload the .env file.
func Get(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
