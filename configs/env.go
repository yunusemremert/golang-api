package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMongoDBURL() string {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	mongoDbUrl := os.Getenv("MONGODB_URL")
	if mongoDbUrl == "" {
		log.Fatal("You must set your 'MONGODB_URL' environmental variable.")
	}

	return mongoDbUrl
}
