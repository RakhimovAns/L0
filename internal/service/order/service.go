package order

import (
	"context"
	"log/slog"

	"github.com/RakhimovAns/L0/internal/model"
	txman "github.com/RakhimovAns/txmananger/tx_manager"
	"github.com/redis/go-redis/v9"
)

type RedisProvider interface {
	GetOrder(ctx context.Context, id string) (*model.Order, error)
}

type Service struct {
	redisClient *redis.Client
	RedisRepo   RedisProvider
	log         *slog.Logger
	tx          txman.TxManager
}

func New(redisClient *redis.Client, redisRepo RedisProvider, log *slog.Logger, tx txman.TxManager) *Service {
	return &Service{
		redisClient: redisClient,
		RedisRepo:   redisRepo,
		log:         log,
		tx:          tx,
	}
}
