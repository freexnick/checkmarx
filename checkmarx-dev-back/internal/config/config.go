package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

const configPath = "./configs/.env"

type Configuration struct {
	ApiVersion uint `env:"HTTP_API_VERSION"`

	HTTPServerAddress   string `env:"HTTP_SERVER_ADDRESS"`
	ReadTimeoutSeconds  uint   `env:"HTTP_SERVER_READ_TIME_OUT"`
	WriteTimeoutSeconds uint   `env:"HTTP_SERVER_WRITE_TIME_OUT"`
	IdleTimeoutSeconds  uint   `env:"HTTP_SERVER_IDLE_TIME_OUT"`

	PostgresURL                  string `env:"POSTGRES_URL"`
	PostgresMinConnections       int    `env:"POSTGRES_MIN_CONNECTIONS"`
	PostgresMaxConnections       int    `env:"POSTGRES_MAX_CONNECTIONS"`
	PostgresMaxIdleTimeoutMinute int    `env:"POSTGRES_MAX_IDLE_TIME_OUT"`
}

func New() (*Configuration, error) {
	conf := &Configuration{}
	if err := godotenv.Load(configPath); err != nil {
		return nil, err
	}

	if err := env.Parse(conf); err != nil {
		return nil, err
	}

	return conf, nil
}
