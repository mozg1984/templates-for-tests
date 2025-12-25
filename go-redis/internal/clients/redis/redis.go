package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func New(address string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         address,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		PoolSize:     10,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd := client.Ping(ctx)

	if err := cmd.Err(); err != nil {
		return nil, fmt.Errorf("server is not responding: %w", err)
	}

	return client, nil
}
