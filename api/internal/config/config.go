package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APP_ENV  string
	APP_PORT string
	APP_URL  string

	DB_HOST    string
	DB_USER    string
	DB_PASS    string
	DB_PORT    string
	DB_NAME    string
	DBSSL_Mode string

	JWTSecret string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	config := &Config{
		APP_ENV:  os.Getenv("APP_ENV"),
		APP_PORT: os.Getenv("APP_PORT"),
		APP_URL:  os.Getenv("APP_URL"),

		DB_HOST:    os.Getenv("DB_HOST"),
		DB_USER:    os.Getenv("DB_USER"),
		DB_PASS:    os.Getenv("DB_PASS"),
		DB_PORT:    os.Getenv("DB_PORT"),
		DB_NAME:    os.Getenv("DB_NAME"),
		DBSSL_Mode: os.Getenv("DBSSL_Mode"),

		JWTSecret: os.Getenv("JWTSecret"),
	}
	if err := validateConfig(config); err != nil {
		return nil, err
	}
	return config, nil
}
func validateConfig(config *Config) error {
	if config.APP_ENV == "production" {
		if config.DB_NAME == "" {
			return fmt.Errorf("DB_NAME is required in production environment")
		}
		if config.DB_PASS == "" {
			return fmt.Errorf("DBPass is required in production environment")
		}
		if config.JWTSecret == "" {
			return fmt.Errorf("JWTSecret is required in production environment")
		}
		if len(config.JWTSecret) < 32 {
			return fmt.Errorf("JWTSecret should be at least 32 characters long for better security")
		}
		if config.DBSSL_Mode == "disable" {
			log.Println("Warning: DBSSLMode is set to 'disable' in production environment. Consider enabling SSL for better security.")
		}
	}
	return nil
}
