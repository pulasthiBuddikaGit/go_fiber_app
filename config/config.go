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
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Failed to load .env file:", err)
	}

	return &Config{
		MongoURI: os.Getenv("MongoDB_URL"),
		Database: os.Getenv("MongoDB_DB"),
	}
}
