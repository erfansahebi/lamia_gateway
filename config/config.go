package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server struct {
		HTTP struct {
			Host string `env:"HOST"`
			Port string `env:"PORT"`
		}
	}
	Services struct {
		Auth struct {
			Host string `env:"LAMIA_AUTH_HOST"`
			Port string `env:"LAMIA_AUTH_PORT"`
		}
	}
}

func LoadConfig() (*Config, error) {
	var configuration Config

	if err := cleanenv.ReadEnv(&configuration); err != nil {
		return nil, err
	}

	return &configuration, nil
}
