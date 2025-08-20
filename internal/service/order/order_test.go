package order_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/RakhimovAns/L0/internal/model"
	"github.com/RakhimovAns/L0/internal/service/order"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"

	txman "github.com/RakhimovAns/txmananger/tx_manager"
)

// мок для RedisProvider
type mockRedisRepo struct{}

func (m *mockRedisRepo) GetOrder(ctx context.Context, id string) (*model.Order, error) {
	return nil, nil
}

// мок для TxManager
type mockTxManager struct{}

func (m *mockTxManager) ReadCommitted(ctx context.Context, h txman.Handler, opts ...txman.TxOption) error {
	return nil
}
func (m *mockTxManager) RepeatableRead(ctx context.Context, h txman.Handler, opts ...txman.TxOption) error {
	return nil
}
func (m *mockTxManager) Serializable(ctx context.Context, h txman.Handler, opts ...txman.TxOption) error {
	return nil
}
func (m *mockTxManager) RunWithOpts(ctx context.Context, h txman.Handler, opts []txman.TxOption) error {
	return nil
}

func TestNewService(t *testing.T) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	repo := &mockRedisRepo{}
	logger := slog.Default()
	tx := &mockTxManager{}

	svc := order.New(redisClient, repo, logger, tx)

	assert.NotNil(t, svc)

}
