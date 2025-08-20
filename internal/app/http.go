package app

import (
	"context"
	"log/slog"

	"github.com/RakhimovAns/L0/internal/config"
	binding "github.com/RakhimovAns/L0/pkg/bindig"
	"github.com/RakhimovAns/wrapper/pkg/closer"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func (a *App) runHttp(ctx context.Context) error {
	app := a.di.HttpServer(ctx)

	log := a.di.Log(ctx).With(
		slog.Group("http",
			slog.String("addr", config.AddrHTTP()),
		),
	)

	closer.Add(func() error {
		log.InfoContext(ctx, "shutting down http server")

		return app.Shutdown()
	})

	a.setupApiHttpHandlers(ctx, app)

	app.Get("/*", static.New("./static/index.html"))

	app.RegisterCustomBinder(binding.NewQueryBinder())
	log.InfoContext(ctx, "go http server!")

	return app.Listen(config.AddrHTTP(), fiber.ListenConfig{
		DisableStartupMessage: true,
	})
}

func (a *App) setupApiHttpHandlers(ctx context.Context, app *fiber.App) {
	api := app.Group("/")
	order := a.di.HttpOrderHandler(ctx)
	order.Setup(api)
}
