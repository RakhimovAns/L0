package di

import (
	"context"

	httporder "github.com/RakhimovAns/L0/internal/controller/order"
	pgorderrepo "github.com/RakhimovAns/L0/internal/repo/order/postgres"
	"github.com/RakhimovAns/L0/internal/service/order"
	diut "github.com/RakhimovAns/L0/pkg/di"
)

func (d *DI) HttpOrderHandler(ctx context.Context) *httporder.Handler {
	return diut.Once(ctx, func(ctx context.Context) *httporder.Handler {
		return httporder.New(d.OrderService(ctx))
	})
}

func (d *DI) OrderService(ctx context.Context) *order.Service {
	return diut.Once(ctx, func(ctx context.Context) *order.Service {
		return order.New(
			d.Redis(ctx),
			d.RedisOrderRepo(ctx),
			d.Log(ctx),
			d.TxManager(ctx),
		)
	})
}

func (d *DI) OrderRepo(ctx context.Context) *pgorderrepo.Repo {
	return diut.Once(ctx, func(ctx context.Context) *pgorderrepo.Repo {
		return pgorderrepo.New(d.Postgres(ctx))
	})
}
