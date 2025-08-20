package di

import (
	"context"
	"fmt"

	"github.com/RakhimovAns/L0/internal/config"
	diut "github.com/RakhimovAns/L0/pkg/di"
	"github.com/RakhimovAns/wrapper/pkg/closer"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

func (d *DI) Redis(ctx context.Context) *redis.Client {
	return diut.Once(ctx, func(ctx context.Context) *redis.Client {
		var redisClient *redis.Client

		timeouts := config.RedisTimeouts()

		redisClient = redis.NewClient(&redis.Options{
			Addr:         fmt.Sprintf("%s:%s", config.RedisHost(), config.RedisPort()),
			Password:     config.RedisPassword(),
			DB:           config.RedisDB(),
			DialTimeout:  timeouts.Connect,
			ReadTimeout:  timeouts.Read,
			WriteTimeout: timeouts.Write,
		})

		_, err := redisClient.Ping(ctx).Result()
		if err != nil {
			d.mustExit(fmt.Errorf("failed to connect to Redis: %w", err))
		}

		if err := redisotel.InstrumentMetrics(redisClient); err != nil {
			d.mustExit(err)
		}

		closer.Add(func() error {
			d.Log(ctx).InfoContext(ctx, "shutting down redis")

			return redisClient.Close()
		})

		return redisClient
	})
}
