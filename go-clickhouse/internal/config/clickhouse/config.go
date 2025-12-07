package clickhouse

import (
	"github.com/caarlos0/env/v11"
)

type Config struct {
	DBName    string `env:"CLICKHOUSE_DATABASE,required"`
	User      string `env:"CLICKHOUSE_USERNAME,required"`
	Password  string `env:"CLICKHOUSE_PASSWORD,required"`
	Endpoints string `env:"CLICKHOUSE_ENDPOINTS,required"`
}

func NewConfig() (cfg Config, err error) {
	err = env.Parse(&cfg)

	return
}
