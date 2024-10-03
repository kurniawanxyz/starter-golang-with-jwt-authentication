package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	AppPort string

	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string

	JWTSecret string
	APIKey    string

	SMTPHost  string
	SMTPPort  string
	SMTPUser  string
	SMTPPass  string
}

var ENV *Config



func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ENV = &Config{
		AppName:   os.Getenv("APP_NAME"),
		AppPort:   os.Getenv("APP_PORT"),
		DBHost:    os.Getenv("DB_HOST"),
		DBPort:    os.Getenv("DB_PORT"),
		DBUser:    os.Getenv("DB_USER"),
		DBPass:    os.Getenv("DB_PASS"),
		DBName:    os.Getenv("DB_NAME"),
		JWTSecret: os.Getenv("JWT_SEC"),
		APIKey:    os.Getenv("API_KEY"),
		SMTPHost:  os.Getenv("SMTP_HOST"),
		SMTPPort:  os.Getenv("SMTP_PORT"),
		SMTPUser:  os.Getenv("SMTP_USER"),
		SMTPPass:  os.Getenv("SMTP_PASS"),
	}
}