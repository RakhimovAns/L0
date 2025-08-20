package di

import (
	"context"

	"github.com/RakhimovAns/L0/internal/config"
	"github.com/RakhimovAns/L0/internal/exerr"
	diut "github.com/RakhimovAns/L0/pkg/di"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
)

func (d *DI) HttpServer(ctx context.Context) *fiber.App {
	return diut.Once(ctx, func(ctx context.Context) *fiber.App {
		timeouts := config.HTTPTimeouts()

		return fiber.New(fiber.Config{
			JSONEncoder:  sonic.Marshal,
			JSONDecoder:  sonic.Unmarshal,
			ErrorHandler: exerr.SendHTTP,
			ReadTimeout:  timeouts.Read,
			WriteTimeout: timeouts.Write,
			IdleTimeout:  timeouts.Idle,
		})
	})
}
