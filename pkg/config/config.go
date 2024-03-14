package config

import "os"

type Config struct {
	PinataApiKey       string `env:"PINATA_API_KEY"`
	PinataSecretApiKey string `env:"PINATA_SECRET_API_KEY"`
	Port               string `env:"PORT"`
}

func LoadConfig() *Config {
	return &Config{
		PinataApiKey:       os.Getenv("PINATA_API_KEY"),
		PinataSecretApiKey: os.Getenv("PINATA_SECRET_API_KEY"),
		Port:               os.Getenv("PORT"),
	}
}
