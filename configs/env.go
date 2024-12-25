package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() []string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env file.")
	}

	mongo_url := os.Getenv("MONGO_URL")
	port := os.Getenv("PORT")

	env := []string{mongo_url, port}

	return env
}
