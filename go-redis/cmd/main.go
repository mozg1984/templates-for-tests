package main

import (
	"context"

	redisClient "go-redis/internal/clients/redis"
	"go-redis/internal/config"
	"go-redis/internal/storage"
	"go-redis/internal/usecase"

	"github.com/sirupsen/logrus"
)

func main() {
	redisClientConfig, err := config.RedisClientConfig()
	if err != nil {
		logrus.Fatal(err)
	}

	client, err := redisClient.New(redisClientConfig.ADDRESS)
	if err != nil {
		logrus.Fatal(err)
	}

	ctx := context.Background()

	cache := storage.NewRepo(client)
	useCase := usecase.New(cache)

	if err = useCase.TestSetAndGet(ctx); err != nil {
		logrus.Fatal(err)
	}

	defer func() {
		if err := client.Conn().Close(); err != nil {
			logrus.Fatal(err)
		}
	}()
}
