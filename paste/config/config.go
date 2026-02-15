package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
	Endpoint         string
	Region           string
	AccessKey        string
	SecretKey        string
	Bucket           string
}

func InitConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found", err)
	}

	return &Config{
		PostgresUser:     os.Getenv("PG_USER"),
		PostgresPassword: os.Getenv("PG_PASSWORD"),
		PostgresDBName:   os.Getenv("PG_NAME"),
		Endpoint:         os.Getenv("S3_ENDPOINT"),
		Region:           os.Getenv("S3_REGION"),
		AccessKey:        os.Getenv("S3_ACCESS_KEY"),
		SecretKey:        os.Getenv("S3_SECRET_KEY"),
		Bucket:           os.Getenv("S3_BUCKET"),
	}
}
