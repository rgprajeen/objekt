package postgres

import (
	"context"

	"github.com/attoleap/objekt/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/stephenafamo/bob"
)

type DB struct {
	DB bob.DB
}

func NewDB(ctx context.Context) (*DB, error) {
	connectionURL, err := config.Get().DB.ConnectionURL()
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.New(ctx, connectionURL)
	if err != nil {
		return nil, err
	}
	stdPool := stdlib.OpenDBFromPool(pool)

	return &DB{DB: bob.NewDB(stdPool)}, nil
}
