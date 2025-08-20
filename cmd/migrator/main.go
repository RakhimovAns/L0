package main

import (
	"context"
	"log/slog"

	di2 "github.com/RakhimovAns/L0/internal/app/di"
	"github.com/RakhimovAns/L0/internal/config"
	"github.com/RakhimovAns/logger/pkg/logger/sl"
	"github.com/RakhimovAns/txmananger/postgres"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	di := di2.DI{}

	db := di.Postgres(ctx)

	log := di.Log(ctx)

	migrator, err := postgres.NewMigrator(db.Pool(), config.PostgresMigrationsPath())
	if err != nil {
		log.Error("failed to setup migrator", sl.ErrAttr(err))

		return
	}

	log.Info("upping the migrations")

	upped, err := migrator.Up(ctx)
	if err != nil {
		log.Error("failed to up migrations", sl.ErrAttr(err))

		return
	}

	if len(upped) == 0 {
		log.Info("there is no migrations to up")

		return
	}

	for _, migration := range upped {
		log.Info("migration upped!", slog.String("name", migration.Source.Path))
	}
}
