package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	Host		string
	User		string
	Password	string
	Name		string
	Port		string
	SSLMode		string
}

func LoadConfig() Config{
	_ = godotenv.Load()

	return Config{
		Host:		os.Getenv("DB_HOST"),
		User:		os.Getenv("DB_USER"),
		Password:	os.Getenv("DB_PASSWORD"),
		Name:		os.Getenv("DB_NAME"),
		Port:		os.Getenv("DB_PORT"),
		SSLMode:	os.Getenv("DB_SSLMODE"),
	}

}