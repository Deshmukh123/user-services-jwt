package config

import (
	"log"
	"github.com/joho/godotenv"
)
// Name     string `json:"name"`

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
