package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
}

func InitConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("no .env file found", err)
	}
}
