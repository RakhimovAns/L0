package redisorderrepo

import (
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	orderCacheKeyPrefix = "order:"
	defaultTTL          = 24 * time.Hour
)

type Repo struct {
	redisClient *redis.Client
}

func New(redisClient *redis.Client) *Repo {
	return &Repo{
		redisClient: redisClient,
	}
}

func orderCacheKey(id string) string {
	return orderCacheKeyPrefix + id
}
