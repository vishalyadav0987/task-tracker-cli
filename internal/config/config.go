package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	DBPath  string
}

func MustLoad() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found")
	}

	return &Config{
		AppPort: os.Getenv("APP_PORT"),
		DBPath:  os.Getenv("DB_PATH"),
	}
}
