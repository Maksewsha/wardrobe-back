package cfg

import (
	"errors"
	"os"
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
	pgUsername, hasEnv := os.LookupEnv("PG_USERNAME")
	if hasEnv {
		return nil, errors.New("PG_USERNAME is not set")
	}
	pgPassword, hasEnv := os.LookupEnv("PG_PASSWORD")
	if hasEnv {
		return nil, errors.New("PG_PASSWORD is not set")
	}
	pgHost, hasEnv := os.LookupEnv("PG_HOST")
	if hasEnv {
		return nil, errors.New("PG_HOST is not set")
	}
	pgPort, hasEnv := os.LookupEnv("PG_PORT")
	if hasEnv {
		return nil, errors.New("PG_PORT is not set")
	}
	pgDBName, hasEnv := os.LookupEnv("PG_DBNAME")
	if hasEnv {
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
