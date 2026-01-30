package api

import (
	"context"
	"fmt"

	"github.com/bd986650/rate-limiter/config"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RateLimiter struct {
	rdb *redis.Client
}

func NewRateLimiter(rdb *redis.Client) *RateLimiter {
	return &RateLimiter{rdb: rdb}
}

func (rl *RateLimiter) Allow(userID string) (bool, error) {
	key := fmt.Sprintf("rate:%s", userID)
	count, err := rl.rdb.Incr(ctx, key).Result()
	if err != nil {
		return false, err
	}
	if count == 1 {
		rl.rdb.Expire(ctx, key, config.RateWindow)
	}
	return count <= int64(config.RateLimit), nil
}
