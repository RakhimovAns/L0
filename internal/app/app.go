package app

import (
	"context"

	"github.com/RakhimovAns/L0/internal/app/di"
	"github.com/RakhimovAns/wrapper/pkg/closer"
	"golang.org/x/sync/errgroup"
)

type App struct {
	di *di.DI
}

func New() *App {
	return &App{
		di: di.New(),
	}
}

func (a *App) Run(ctx context.Context) error {
	wg, ctx := errgroup.WithContext(ctx)

	closer.SetLogger(a.di.Log(ctx))

	if err := a.WarmupCaches(ctx); err != nil {
		return err
	}

	wg.Go(func() error {
		return a.runHttp(context.Background())
	})
	wg.Go(func() error {
		return a.RunKafka(context.Background())
	})

	return wg.Wait()
}
