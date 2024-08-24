package config

import "github.com/joho/godotenv"

func envLoad() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}
