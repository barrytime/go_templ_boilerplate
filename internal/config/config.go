package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Env string

const (
	EnvTest Env = "test"
	EnvDev  Env = "dev"
	EnvProd Env = "prod"
)

type Config struct {
	Env              Env    `env:"ENV"`
	ApiServerAddr    string `env:"API_SERVER_ADDR"`
	ApiServerPort    string `env:"API_SERVER_PORT"`
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresPort     string `env:"POSTGRES_PORT"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresDB       string `env:"POSTGRES_DB"`
	PostgresPortTest string `env:"TEST_POSTGRES_PORT"`
	RedisHost        string `env:"REDIS_HOST"`
	RedisPort        string `env:"REDIS_PORT"`
	RedisDB          string `env:"REDIS_DB"`
	SessionName      string `env:"SESSION_NAME"`
	SessionSecret    string `env:"SESSION_SECRET"`
	SqlPath          string `env:"SQL_PATH"`
}

func (c Config) DatabaseURL() string {
	port := c.PostgresPort
	if c.Env == EnvTest {
		port = c.PostgresPortTest
	}
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", c.PostgresUser, c.PostgresPassword, c.PostgresHost, port, c.PostgresDB)
}

func (c Config) RedisURL() string {
	return fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort)
}

func New() (*Config, error) {
	conf, err := env.ParseAs[Config]()
	if err != nil {
		return nil, fmt.Errorf("failed to load config, %w", err)
	}
	return &conf, nil
}

func (c Config) AddressString() (string, error) {
	if c.ApiServerAddr == "" && c.ApiServerPort == "" {
		return "", fmt.Errorf("error API_SERVER_ADDR and API_SERVER_PORT must be set")
	}
	return fmt.Sprintf("%s:%s", c.ApiServerAddr, c.ApiServerPort), nil
}
