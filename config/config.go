package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken string
	DatabaseURL   string
	Debug         bool
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		TelegramToken: getEnv("TELEGRAM_BOT_TOKEN", ""),
		DatabaseURL:   getEnv("DATABASE_URL", ""),
		Debug:         getEnv("DEBUG", "false") == "false",
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
