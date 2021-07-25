package config

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port int `env:"PORT" env-default:"8000"`
}

func Read() (*Config, error) {
	var config Config
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		return nil, err
	}
	if os.Getenv("IS_DEV") != "" {
		err = cleanenv.ReadConfig(".env", &config)
	}

	return &config, err
}
