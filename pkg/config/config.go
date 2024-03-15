package config

import (
	"os"
	"strings"
)

type Config struct {
	PinataApiKey       string `env:"PINATA_API_KEY"`
	PinataSecretApiKey string `env:"PINATA_SECRET_API_KEY"`
	Port               string `env:"PORT"`
}

func LoadConfig() *Config {
	return &Config{
		PinataApiKey:       strings.TrimSpace(os.Getenv("PINATA_API_KEY")),
		PinataSecretApiKey: strings.TrimSpace(os.Getenv("PINATA_SECRET_API_KEY")),
		Port:               strings.TrimSpace(os.Getenv("PORT")),
	}
}
