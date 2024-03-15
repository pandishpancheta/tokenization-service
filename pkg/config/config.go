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
		PinataApiKey:       strings.TrimSuffix(os.Getenv("PINATA_API_KEY"), "\n"),
		PinataSecretApiKey: strings.TrimSuffix(os.Getenv("PINATA_SECRET_API_KEY"), "\n"),
		Port:               strings.TrimSuffix(os.Getenv("PORT"), "\n"),
	}
}
