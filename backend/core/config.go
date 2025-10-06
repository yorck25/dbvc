package core

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JwtSecretkey []byte
}

func LoadConfig() (*Config, error) {
	godotenv.Load()

	config := &Config{}

	key := os.Getenv("SECRET_KEY")
	if key == "" {
		return nil, errors.New("no secret key")
	}
	config.JwtSecretkey = []byte(key)

	return config, nil
}
