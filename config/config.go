package config

type Config struct {
	MongoURI string
	Database string
}

func LoadConfig() *Config {
	return &Config{
		MongoURI: "mongodb://localhost:27017", // or 127.0.0.1:27017
		Database: "userDB",
	}
}
