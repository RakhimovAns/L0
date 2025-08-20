package di

import (
	"context"

	redisorderrepo "github.com/RakhimovAns/L0/internal/repo/order/redis"
	diut "github.com/RakhimovAns/L0/pkg/di"
)

func (d *DI) RedisOrderRepo(ctx context.Context) *redisorderrepo.Repo {
	return diut.Once(ctx, func(ctx context.Context) *redisorderrepo.Repo {
		return redisorderrepo.New(d.Redis(ctx))
	})
}
