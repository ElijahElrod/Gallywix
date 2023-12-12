// Package config holds the configuration objects for the crypto-exchange, storage database, and logger configuration
package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Exchange ExchangeConfig `env:",prefix=EXCHANGE_,required"`
	Database DatabaseConfig `env:",prefix=DB_,required"`
	Logger   LoggerConfig   `env:",prefix=LOGGER_"`
}

// ExchangeConfig holds the configuration & access details for exchange
type ExchangeConfig struct {
	Wss              string   `env:"WSS,required"`
	Url              string   `env:"URL,required"`
	Origin           string   `env:"ORIGIN,required"`
	Protocol         string   `env:"PROTOCOL,default="`
	Symbols          []string `env:"SYMBOLS,required"`
	Channels         []string `env:"CHANNELS,required"`
	AccessKey        string   `env:"ACCESS_KEY,required"`
	AccessPassphrase string   `env:"ACCESS_PASSPHRASE,required"`
	AccessSecret     string   `env:"ACCESS_SECRET,required"`
}

// DatabaseConfig holds the configuration & access details for the database
type DatabaseConfig struct {
	Host     string `env:"HOST,required"`
	User     string `env:"USER,required"`
	Password string `env:"PASSWORD"`
	Base     string `env:"BASE"`
}

// LoggerConfig holds the configuration & access details for the zap logger
type LoggerConfig struct {
	DisableCaller     bool   `env:"CALLER,default=false"`
	DisableStacktrace bool   `env:"STACKTRACE,default=false"`
	Level             string `env:"LEVEL,default=debug"`
}

// NewConfig creates a new configuration holding ExchangeConfig, DatabaseConfig, and LoggerConfig for Application Startup
func NewConfig(ctx context.Context) (*Config, error) {
	var cfg Config

	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
