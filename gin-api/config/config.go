package config

import (
	"os"

	"github.com/joho/godotenv"
)

var AppConfig struct {
	Port        string
	DatabaseURL string
}

// Load reads configuration from .env file
func Load() error {
	err := godotenv.Load("app.env")
	if err != nil {
		return err
	}

	AppConfig.Port = os.Getenv("PORT")
	AppConfig.DatabaseURL = os.Getenv("DATABASE_URL")
	return nil
}

// GetPort returns the port number
func GetPort() string {
	if AppConfig.Port == "" {
		return "8080" // default port
	}
	return AppConfig.Port
}
