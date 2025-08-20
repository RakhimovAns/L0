package httporder

import (
	"context"

	"github.com/RakhimovAns/L0/internal/model"
	"github.com/RakhimovAns/L0/internal/service/order"
	"github.com/gofiber/fiber/v3"
)

type Service interface {
	FetchByID(ctx context.Context, in order.FetchIn) (model.Order, error)
}

type Handler struct {
	service Service
}

func New(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Setup(router fiber.Router) {
	group := router.Group("/order")
	group.Post("/ping", h.ping)
	group.Get("/:order_uid", h.fetchByID)
}
