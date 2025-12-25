package config

import (
	"github.com/caarlos0/env/v11"
)

type redisClientConfig struct {
	ADDRESS string `env:"REDIS_ADDRESS,required"`
}

func RedisClientConfig() (cfg redisClientConfig, err error) {
	err = env.Parse(&cfg)

	return
}
