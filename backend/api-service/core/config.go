package core

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	JwtSecretkey []byte

	psqlHost     string
	psqlPort     int
	psqlUser     string
	psqlPassword string
	psqlDatabase string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := &Config{}

	key := os.Getenv("SECRET_KEY")
	if key == "" {
		return nil, errors.New("no secret key")
	}
	config.JwtSecretkey = []byte(key)

	psqlHost := os.Getenv("PSQL_HOST")
	if psqlHost == "" {
		return nil, errors.New("no psql host")
	}
	config.psqlHost = psqlHost

	psqlPort, err := strconv.Atoi(os.Getenv("PSQL_PORT"))
	if err != nil {
		return nil, errors.New("no psql port")
	}
	config.psqlPort = psqlPort

	psqlUser := os.Getenv("PSQL_USER")
	if psqlUser == "" {
		return nil, errors.New("no psql user")
	}
	config.psqlUser = psqlUser

	psqlPassword := os.Getenv("PSQL_PASSWORD")
	if psqlPassword == "" {
		return nil, errors.New("no psql password")
	}
	config.psqlPassword = psqlPassword

	psqlDatabase := os.Getenv("PSQL_DATABASE")
	if psqlDatabase == "" {
		return nil, errors.New("no psql database")
	}
	config.psqlDatabase = psqlDatabase

	return config, nil
}
