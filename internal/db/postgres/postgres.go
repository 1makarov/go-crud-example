package postgres

import (
	"github.com/1makarov/go-crud-example/internal/db"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func Open(cfg db.ConfigDB) (*sqlx.DB, error) {
	client, err := sqlx.Connect("pgx", cfg.String())
	if err != nil {
		return nil, err
	}

	return client, nil
}
