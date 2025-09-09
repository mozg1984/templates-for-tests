package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type DBConfig struct {
	Host     string `env:"DB_HOST,required"`
	Port     string `env:"DB_PORT,required"`
	DBName   string `env:"DB_NAME,required"`
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	SSLMode  string `env:"DB_SSL_MODE" envDefault:"disable"`

	MaxConnections           int32 `env:"DB_MAX_CONNECTIONS" envDefault:"5"`
	MaxConnectionLifetimeSec int   `env:"DB_MAX_CONNECTION_LIFETIME_SEC" envDefault:"120"` // default 2 min
	MaxConnectionIdleTimeSec int   `env:"DB_MAX_CONNECTION_IDLE_TIME_SEC" envDefault:"30"`

	EnableLogging bool `env:"ENABLE_LOGGING" envDefault:"false"`
}

func NewDBConfig() (cfg DBConfig, err error) {
	err = env.Parse(&cfg)

	return
}

func (c DBConfig) ConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.DBName, c.SSLMode)
}
