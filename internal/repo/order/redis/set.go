package redisorderrepo

import (
	"context"
	"encoding/json"

	"github.com/RakhimovAns/L0/internal/model"
	slerr "github.com/RakhimovAns/logger/pkg/err"
)

func (r *Repo) SetOrder(ctx context.Context, order model.Order) error {
	data, err := json.Marshal(order)
	if err != nil {
		return slerr.WithSource(err)
	}

	key := orderCacheKey(order.OrderUID)

	if err := r.redisClient.Set(ctx, key, data, defaultTTL).Err(); err != nil {
		return slerr.WithSource(err)
	}

	return nil
}
