package app

import (
	"context"
	"fmt"

	slerr "github.com/RakhimovAns/logger/pkg/err"
	"github.com/RakhimovAns/wrapper/pkg/closer"
)

func (a *App) WarmupCaches(ctx context.Context) error {
	a.di.Log(ctx).Info("Starting cache warmup...")
	closer.Add(func() error {
		a.di.Log(ctx).Info("shutting init worker...")
		return nil
	})

	orderRepo := a.di.OrderRepo(ctx)
	orderRedis := a.di.RedisOrderRepo(ctx)

	ids, err := orderRepo.FetchAllIDs(ctx)
	if err != nil {
		return slerr.WithSource(err)
	}

	for _, id := range ids {
		newOrder, err := orderRepo.FetchByID(ctx, id)
		if err != nil {
			fmt.Print("eror", err.Error())
			a.di.Log(ctx).Error("failed to fetch order", "order_uid", id, "error", err.Error())
			continue
		}

		if err := orderRedis.SetOrder(ctx, newOrder); err != nil {
			a.di.Log(ctx).Error("failed to set order in Redis", "order_uid", id, "error", err)
		}
	}

	a.di.Log(ctx).Info("Cache warmup completed")
	a.di.Log(ctx).Info("all orders cached in Redis", "count", len(ids))
	return nil
}
