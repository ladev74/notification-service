package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"

	"notification/internal/logger"
	"notification/internal/notification/api"
	"notification/internal/notification/service"
	"notification/internal/rds"
)

type Config struct {
	HttpServer api.HttpServer `yaml:"HTTP_SERVER"`
	SMTP       service.Config `yaml:"SMTP"`
	Redis      rds.Config     `yaml:"REDIS"`
	Logger     logger.Config  `yaml:"LOGGER"`
}

func New(path string) (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	return &cfg, nil
}
