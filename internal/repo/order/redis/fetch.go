package redisorderrepo

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/RakhimovAns/L0/internal/model"
	slerr "github.com/RakhimovAns/logger/pkg/err"
	"github.com/redis/go-redis/v9"
)

func (r *Repo) GetOrder(ctx context.Context, id string) (*model.Order, error) {
	key := orderCacheKey(id)

	val, err := r.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, slerr.WithSource(err)
	}

	var order model.Order
	if err := json.Unmarshal(val, &order); err != nil {
		return nil, slerr.WithSource(err)
	}

	return &order, nil
}
