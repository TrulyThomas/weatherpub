package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type RabbitMQ struct {
	URI string `yaml:"URI"`
}

type Config struct {
	Rabbit RabbitMQ `yaml:"RabbitMQ"`
}

func NewConfig(path string) (*Config, error) {
	var cfg Config

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
