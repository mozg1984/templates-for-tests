package config

import (
	"github.com/caarlos0/env/v11"
)

type kafkaConfig struct {
	Brokers []string `env:"KAFKA_BOOTSTRAP_SERVERS,required"`
}

func KafkaConfig() (cfg kafkaConfig, err error) {
	err = env.Parse(&cfg)

	return
}
