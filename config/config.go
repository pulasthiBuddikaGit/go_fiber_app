package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	Database string
}

func LoadConfig() *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file:", err)
	}

	return &Config{
		MongoURI: os.Getenv("MongoDB_URL"),
		Database: os.Getenv("MongoDB_DB"),
	}
}
