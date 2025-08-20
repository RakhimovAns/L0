package order

import (
	"context"

	"github.com/RakhimovAns/L0/internal/exerr"
	"github.com/RakhimovAns/L0/internal/model"
	slerr "github.com/RakhimovAns/logger/pkg/err"
	txman "github.com/RakhimovAns/txmananger/tx_manager"
)

type FetchIn struct {
	ID string `json:"id"`
}

func (s *Service) FetchByID(ctx context.Context, in FetchIn) (model.Order, error) {
	var order *model.Order
	var err error
	
	err = s.tx.Serializable(ctx, func(ctx context.Context) error {
		order, err = s.RedisRepo.GetOrder(ctx, in.ID)
		if err != nil {
			return slerr.WithSource(err)
		}

		if order == nil {
			return slerr.WithSource(exerr.New("order not found"))
		}

		return nil
	}, txman.WithRetry(1))
	if err != nil {
		return model.Order{}, err
	}

	return *order, nil
}
