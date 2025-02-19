package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BOTTOKEN      string
	APP_ENV       string
	OPEN_AI_TOKEN string
}

func LoadConfig() (*Config, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Determine the environment
	appEnv := os.Getenv("APP_ENV")

	// OpenAI Key

	openAiToken := os.Getenv("OPEN_AI_TOKEN")

	// Default to test bot token
	botToken := os.Getenv("TEST_BOT_TOKEN")

	// If running in production, override with production bot token
	if appEnv == "PROD" {
		prodToken := os.Getenv("BOT_TOKEN")
		if prodToken != "" {
			botToken = prodToken
		}
	}

	// Ensure that we actually have a bot token
	if botToken == "" {
		return nil, fmt.Errorf("missing required environment variables: BOT_TOKEN")
	}

	// Ensure that we actually have a bot token
	if openAiToken == "" {
		return nil, fmt.Errorf("missing required environment variables: OPEN_AI_TOKEN")
	}

	return &Config{
		BOTTOKEN:      botToken,
		APP_ENV:       appEnv,
		OPEN_AI_TOKEN: openAiToken,
	}, nil

}
