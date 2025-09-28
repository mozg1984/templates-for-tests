package config

import (
	"github.com/caarlos0/env/v11"
)

type mongoDBConfig struct {
	URI string `env:"MONGODB_URI,required"`
}

func MongoDBConfig() (cfg mongoDBConfig, err error) {
	err = env.Parse(&cfg)

	return
}
