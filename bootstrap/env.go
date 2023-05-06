package bootstrap

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	DB_HOST string
	DB_USER string
	DB_PASSWORD string
	DB_DATABASE string
	DB_PORT string
	DB_DRIVER string
	APP_URL string
}

func NewEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error message is => ", err)
	}
	
	return &Env{
		DB_HOST: os.Getenv("DB_HOST"),
		DB_USER: os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_DATABASE: os.Getenv("DB_DATABASE"),
		DB_PORT: os.Getenv("DB_PORT"),
		DB_DRIVER: os.Getenv("DB_DRIVER"),
		APP_URL: os.Getenv("APP_URL"),
	}
}
