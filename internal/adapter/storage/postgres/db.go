package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.prajeen.com/objekt/internal/config"
)

type DB struct {
	Pool *pgxpool.Pool
}

func NewDB(ctx context.Context, db *config.DB) (*DB, error) {
	pool, err := pgxpool.New(ctx, db.ConnectionURL())
	if err != nil {
		return nil, err
	}

	return &DB{Pool: pool}, nil
}
