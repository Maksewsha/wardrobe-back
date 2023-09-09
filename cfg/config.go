package cfg

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type PGConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

type Config struct {
	PGConfig *PGConfig
}

func New() (*Config, error) {
	cfg := &Config{}
	err := godotenv.Load(".env")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to retrieve config values from .env: %v\n", err))
	}
	pgUsername, hasEnv := os.LookupEnv("PG_USERNAME")
	if !hasEnv {
		return nil, errors.New("PG_USERNAME is not set")
	}
	pgPassword, hasEnv := os.LookupEnv("PG_PASSWORD")
	if !hasEnv {
		return nil, errors.New("PG_PASSWORD is not set")
	}
	pgHost, hasEnv := os.LookupEnv("PG_HOST")
	if !hasEnv {
		return nil, errors.New("PG_HOST is not set")
	}
	pgPort, hasEnv := os.LookupEnv("PG_PORT")
	if !hasEnv {
		return nil, errors.New("PG_PORT is not set")
	}
	pgDBName, hasEnv := os.LookupEnv("PG_DBNAME")
	if !hasEnv {
		return nil, errors.New("PG_DBNAME is not set")
	}
	cfg.PGConfig = &PGConfig{
		Username: pgUsername,
		Password: pgPassword,
		Host:     pgHost,
		Port:     pgPort,
		DBName:   pgDBName,
	}
	return cfg, nil
}
