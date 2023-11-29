package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env File")
		return err
	}
	return nil
}
