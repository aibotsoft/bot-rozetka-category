package config

import (
	"github.com/caarlos0/env/v6"
	"time"
)

type Config struct {
	ApiDebug    bool   `env:"API_DEBUG"`
	PostgresDSN string `env:"POSTGRES_DSN,required"`
	// debug, info, warn, error, fatal, panic
	LogLever string `env:"LOG_LEVEL" envDefault:"info"`
	// json or console
	LogEncoding    string        `env:"LOG_ENCODING" envDefault:"json"`
	RootCategoryID int64         `env:"ROOT_CATEGORY_ID" envDefault:"83850"`
	Port           string        `env:"PORT" envDefault:"8080"`
	RequestTimeout time.Duration `env:"REQUEST_TIMEOUT" envDefault:"5s"`
	TelegramToken  string        `env:"TELEGRAM_TOKEN"`
	TelegramDebug  bool          `env:"TELEGRAM_DEBUG"`
	Version        string        `env:"VERSION"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	return cfg, err
}
