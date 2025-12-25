package storage

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Repo struct {
	redisCli *redis.Client
}

func NewRepo(client *redis.Client) *Repo {
	return &Repo{
		redisCli: client,
	}
}

func (r *Repo) Get(ctx context.Context, key string) (val string, err error) {
	strCmd := r.redisCli.Get(ctx, key)

	if err = strCmd.Err(); err != nil {
		return
	}

	val = strCmd.Val()

	return
}

func (r *Repo) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	return r.redisCli.Set(ctx, key, value, ttl).Err()
}
