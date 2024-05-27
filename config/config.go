package config

import (
	"sync"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type ServiceConfig struct {
	ID  string `env:"ID"`
	Key string `env:"KEY"`
}

type Config struct {
	Env     string        `env:"ENV"`
	Service ServiceConfig `envPrefix:"SERVICE_"`
}

var (
	Cfg        *Config
	configOnce sync.Once
)

func parseConfig() (*Config, error) {
	// https://github.com/spf13/viper use this if complicated config is required
	// makes all fields required if default is not defined
	err := godotenv.Load("/Users/ayush/simpl/golang-learning/development.env")
	if err != nil {
		return nil, err
	}

	opts := env.Options{RequiredIfNoDef: true}

	cfg := &Config{}
	if err := env.Parse(cfg, opts); err != nil {
		return nil, err
	}
	return cfg, nil
}

func NewConfig() *Config {
	configOnce.Do(func() {
		var err error
		Cfg, err = parseConfig()
		if err != nil {
			panic(err)
		}
	})

	return Cfg
}
