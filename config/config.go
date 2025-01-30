package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BOTTOKEN string
}

func LoadConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botToken := os.Getenv("BOT_TOKEN")

	if botToken == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	return &Config{
		BOTTOKEN: botToken,
	}, nil

}
