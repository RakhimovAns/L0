package httporder

import (
	"github.com/RakhimovAns/L0/internal/service/order"
	slerr "github.com/RakhimovAns/logger/pkg/err"
	"github.com/gofiber/fiber/v3"
)

func (h *Handler) ping(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "pong",
	})
}

func (h *Handler) fetchByID(ctx fiber.Ctx) error {
	orderUID := ctx.Params("order_uid")

	if orderUID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "order_uid is required",
		})
	}

	order, err := h.service.FetchByID(ctx, order.FetchIn{ID: orderUID})
	if err != nil {
		return slerr.WithSource(err)
	}

	return ctx.JSON(order)
}
