package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/stephenafamo/bob"
)

type DB struct {
	Pool *pgxpool.Pool
	DB   bob.DB
}

func NewDB(ctx context.Context, connectionURL string) (*DB, error) {
	pool, err := pgxpool.New(ctx, connectionURL)
	if err != nil {
		return nil, err
	}
	stdPool := stdlib.OpenDBFromPool(pool)

	return &DB{Pool: pool, DB: bob.NewDB(stdPool)}, nil
}
