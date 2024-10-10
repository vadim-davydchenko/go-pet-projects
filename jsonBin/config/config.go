package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error for download .env file: %v", err)
	}

	key := os.Getenv("KEY")
	if key == "" {
		return nil, fmt.Errorf("variable environment KEY don't found")
	}
	return &Config{
		Key: key,
	}, nil
}
