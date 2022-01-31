package config

import (
	"gopkg.in/yaml.v2"
	"os"
)

type App struct {
	ServicePort string `yaml:"service_port"`
	HandlerMode string `yaml:"handler_mode"`
}

type Elasticsearch struct {
	Address string `yaml:"es_address"`
}

type Config struct {
	App
	Elasticsearch
}

func (c *Config) Init() (*Config, error) {
	file, err := os.Open("config/config.yml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := yaml.NewDecoder(file)

	if err = decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, err
}
