package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
	JWTSecret   string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file: ", err)
	}

	config := &Config{
		DatabaseURL: os.Getenv("DATABASE"),
		Port:        os.Getenv("PORT"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}

	return config, err
}
