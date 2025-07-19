package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppPort       string
	AppHost       string
	DbHost        string
	DbUser        string
	DbPassword    string
	DbName        string
	DbPort        string
	DbSSLMode     string
	EmailHost     string
	EmailPort     string
	EmailUsername string
	EmailPassword string
}



func NewConfigInstance(envFile string) *AppConfig {
	if err := godotenv.Load(envFile); err != nil {
		panic("Error loading env file: " + err.Error())
	}

	return &AppConfig{
		AppPort:       os.Getenv("PORT"),
		AppHost:       os.Getenv("APP_HOST"),
		DbHost:        os.Getenv("DB_HOST"),
		DbUser:        os.Getenv("DB_USER"),
		DbPassword:    os.Getenv("DB_PASSWORD"),
		DbName:        os.Getenv("DB_NAME"),
		DbPort:        os.Getenv("DB_PORT"),
		DbSSLMode:     os.Getenv("DB_SSL"),
		EmailHost:     os.Getenv("EMAIL_HOST"),
		EmailPort:     os.Getenv("EMAIL_PORT"),
		EmailUsername: os.Getenv("EMAIL_USERNAME"),
		EmailPassword: os.Getenv("EMAIL_PASSWORD"),
	}
}
