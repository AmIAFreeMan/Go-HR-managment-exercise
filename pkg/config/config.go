package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL   string
	Auth0Issuer   string
	Auth0Audience string
	ServerPort    string
}

func Load() (*Config, error) {
	// Load .env file if present
	_ = godotenv.Load()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":50051"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/hr_system?sslmode=disable"
	}

	return &Config{
		DatabaseURL:   dbURL,
		Auth0Issuer:   os.Getenv("AUTH0_ISSUER_URL"),
		Auth0Audience: os.Getenv("AUTH0_AUDIENCE"),
		ServerPort:    port,
	}, nil
}
