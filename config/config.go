// Package config
package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

var envs Config

type AppEnvironment struct {
	BasePath string `env:"BASE_PATH"`
	AppHost  string `env:"APP_DEFAULT_HOST" envDefault:"0.0.0.0"`
	AppPort  string `env:"APP_DEFAULT_PORT" envDefault:"8080"`
}

type DatabaseEnvironment struct {
	DatabaseName     string `env:"DB_NAME"`
	DatabaseHost     string `env:"DB_HOST"`
	DatabasePort     string `env:"DB_PORT"`
	DatabaseUser     string `env:"DB_USER"`
	DatabasePassword string `env:"DB_PASSWORD"`
}

type Config struct {
	AppEnvironment
	DatabaseEnvironment
}

func ParseEnv() error {
	err := env.Parse(&envs)
	if err != nil {
		return fmt.Errorf("failed to load config %s", err)
	}
	return nil
}

func (c *Config) DatabaseURL() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		c.DatabaseUser,
		c.DatabasePassword,
		c.DatabaseHost,
		c.DatabasePort,
		c.DatabaseName,
	)
}

func GetEnv() *Config {
	return &envs
}
