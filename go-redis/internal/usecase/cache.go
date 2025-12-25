package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type cacheRepo interface {
	Get(ctx context.Context, key string) (val string, err error)
	Set(ctx context.Context, key string, value string, ttl time.Duration) error
}

type UseCase struct {
	cache cacheRepo
}

func New(cache cacheRepo) *UseCase {
	return &UseCase{cache: cache}
}

func (u *UseCase) TestSetAndGet(ctx context.Context) error {
	cacheKey := "key:1234567891"

	_, err := u.cache.Get(ctx, cacheKey)
	if err != nil && !errors.Is(err, redis.Nil) {
		return fmt.Errorf("getting from cache error: %w", err)
	}

	err = u.cache.Set(ctx, cacheKey, "1", 5*time.Second)
	if err != nil {
		return fmt.Errorf("setting to cache error: %w", err)
	}

	val, err := u.cache.Get(ctx, cacheKey)
	if err != nil && !errors.Is(err, redis.Nil) {
		return fmt.Errorf("getting from cache error: %w", err)
	}

	fmt.Println("val: ", val)

	time.Sleep(5 * time.Second)

	val, err = u.cache.Get(ctx, cacheKey)
	if err != nil && !errors.Is(err, redis.Nil) {
		return fmt.Errorf("getting from cache error: %w", err)
	}

	fmt.Println("val (after ttl expiration): ", val)

	return nil
}
