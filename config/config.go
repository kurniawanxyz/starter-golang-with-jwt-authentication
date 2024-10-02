package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	AppPort int

	DBHost string
	DBPort int
	DBUser string
	DBPass string
	DBName string

	JWTSecret string
	APIKey    string
}

var ENV *Config



func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	appPort, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		log.Fatalf("Error converting APP_PORT to int: %v", err)
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Error converting DB_PORT to int: %v", err)
	}

	ENV = &Config{
		AppName:   os.Getenv("APP_NAME"),
		AppPort:   appPort,
		DBHost:    os.Getenv("DB_HOST"),
		DBPort:    dbPort,
		DBUser:    os.Getenv("DB_USER"),
		DBPass:    os.Getenv("DB_PASS"),
		DBName:    os.Getenv("DB_NAME"),
		JWTSecret: os.Getenv("JWT_SEC"),
		APIKey:    os.Getenv("API_KEY"),
	}
}