package di

import (
	"context"

	"github.com/RakhimovAns/L0/internal/config"
	diut "github.com/RakhimovAns/L0/pkg/di"
	"github.com/segmentio/kafka-go"
)

func (d *DI) KafkaConsumer(ctx context.Context) *kafka.Reader {
	return diut.Once(ctx, func(ctx context.Context) *kafka.Reader {
		cfg := config.Kafka()

		r := kafka.NewReader(kafka.ReaderConfig{
			Brokers:     []string{cfg.Broker},
			Topic:       cfg.Topic,
			GroupID:     cfg.GroupID,
			StartOffset: kafka.LastOffset,
		})

		return r
	})
}
