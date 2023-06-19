package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
	"path"
)

type Config struct {
	Server struct {
		HTTP struct {
			Host string `env:"HOST"`
			Port string `env:"PORT"`
		}
	}
	JWT struct {
		Secret string `env:"JWT_SECRET"`
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

	configPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	if err = godotenv.Load(path.Join(configPath, ".env")); err != nil {
		return nil, err
	}

	if err = cleanenv.ReadEnv(&configuration); err != nil {
		return nil, err
	}

	return &configuration, nil
}
