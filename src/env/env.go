package env

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func Process(key string) string {
	//err := godotenv.Load()
	err := godotenv.Load(filepath.Join("./env/", ".env"))
	if err != nil {
		log.Fatal("Error loading .env")
	}
	return os.Getenv(key)
}
