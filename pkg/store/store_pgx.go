package pkgstore

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type StorePgx struct {
	DB  *pgxpool.Pool
	Ctx context.Context
}

func NewStorePgx(ctx context.Context, connstring string) (*StorePgx, error) {
	pool, err := pgxpool.New(ctx, connstring)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}
	return &StorePgx{DB: pool, Ctx: ctx}, nil
}
