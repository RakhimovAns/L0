package app

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/RakhimovAns/L0/internal/model"
	slerr "github.com/RakhimovAns/logger/pkg/err"
	txman "github.com/RakhimovAns/txmananger/tx_manager"
	"golang.org/x/sync/errgroup"
)

func (a *App) RunKafka(ctx context.Context) error {
	reader := a.di.KafkaConsumer(ctx)
	orderRepo := a.di.OrderRepo(ctx)
	orderRedis := a.di.RedisOrderRepo(ctx)
	tx := a.di.TxManager(ctx)
	log := a.di.Log(ctx).With(
		slog.String("component", "kafka"),
	)

	log.Info("starting kafka consumer")

	wg, ctx := errgroup.WithContext(ctx)

	wg.Go(func() error {
		for {
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				log.Error("kafka read error", "error", err.Error())
				continue
			}

			var order model.Order
			if err := json.Unmarshal(m.Value, &order); err != nil {
				log.Error("invalid order JSON", "err", err)
				continue
			}
			err = tx.Serializable(ctx, func(ctx context.Context) error {
				rawOrder, err := orderRedis.GetOrder(ctx, order.OrderUID)
				if err != nil {
					log.Error("failed to get order from redis", "err", err)
					return slerr.WithSource(err)
				}

				if rawOrder != nil {
					log.Info("order already exists in redis", "order_uid", order.OrderUID)
					return slerr.WithSource(err)
				}

				if err = orderRepo.Create(ctx, order); err != nil {
					log.Error("failed to create order", "err", err.Error())
					return slerr.WithSource(err)
				}

				log.Info("order created", "order_uid", order.OrderUID)

				if err = orderRedis.SetOrder(ctx, order); err != nil {
					log.Error("failed to set order in redis", "err", err)
					return slerr.WithSource(err)
				}

				return nil
			}, txman.WithRetry(1))
			if err != nil {
				log.Error("failed to serialize order", "err", err)
				continue
			}
			log.Info("order set in redis", "order_uid", order.OrderUID)
		}
	})

	return wg.Wait()
}
