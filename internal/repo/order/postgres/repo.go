package pgorderrepo

import (
	"github.com/RakhimovAns/txmananger/postgres"
	"github.com/huandu/go-sqlbuilder"
)

type Repo struct {
	db postgres.Postgres
	qb sqlbuilder.Flavor
}

func New(db postgres.Postgres) *Repo {
	return &Repo{
		db: db,
		qb: sqlbuilder.PostgreSQL,
	}
}
