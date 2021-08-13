package config

import "github.com/caarlos0/env"

type Config struct {
	LogLevel           string `env:"LOG_LEVEL" envDefault:"debug"`
	ItemLimit          int    `env:"ITEM_LIMIT" envDefault:"40"`
	RequestInterval    int    `env:"REQUEST_INTERVAL" envDefault:"1"`
	ParseIntervalHours int    `env:"PARSE_INTERVAL_HOURS" envDefault:"6"`
	MaxWorkers         int    `env:"MAX_WORKERS" envDefault:"3"`
}

func NewConfig() (*Config, error) {
	cfg := new(Config)
	return cfg, env.Parse(cfg)
}
