package core

import (
	"backend/common"
	"errors"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func LoadConfig() (*common.Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	config := &common.Config{}

	key := os.Getenv("SECRET_KEY")
	if key == "" {
		return nil, errors.New("no secret key")
	}
	config.JwtSecretKey = []byte(key)

	psqlHost := os.Getenv("PSQL_HOST")
	if psqlHost == "" {
		return nil, errors.New("no psql host")
	}
	config.PsqlHost = psqlHost

	psqlPort, err := strconv.Atoi(os.Getenv("PSQL_PORT"))
	if err != nil {
		return nil, errors.New("no psql port")
	}
	config.PsqlPort = psqlPort

	psqlUser := os.Getenv("PSQL_USER")
	if psqlUser == "" {
		return nil, errors.New("no psql user")
	}
	config.PsqlUser = psqlUser

	psqlPassword := os.Getenv("PSQL_PASSWORD")
	if psqlPassword == "" {
		return nil, errors.New("no psql password")
	}
	config.PsqlPassword = psqlPassword

	psqlDatabase := os.Getenv("PSQL_DATABASE")
	if psqlDatabase == "" {
		return nil, errors.New("no psql database")
	}
	config.PsqlDatabase = psqlDatabase

	return config, nil
}
