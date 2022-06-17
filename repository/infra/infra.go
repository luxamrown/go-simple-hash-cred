package infra

import (
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type Infra interface {
	Connect() *sqlx.DB
}

type infra struct {
	db *sqlx.DB
}

func (i *infra) Connect() *sqlx.DB {
	return i.db
}

func NewInfra() Infra {
	conn, err := sqlx.Connect("pgx", "postgres://postgres:stauffenberg@localhost:5432/bank?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return &infra{
		db: conn,
	}
}
