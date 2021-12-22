package config

type Config struct {
	ApiDebug    bool   `env:"API_DEBUG,required"`
	PostgresDSN string `env:"POSTGRES_DSN,required"`
}
