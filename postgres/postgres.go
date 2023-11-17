package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	ctx    context.Context
	config pgxpool.Config
}

func New(confString string, ctx context.Context) (*Postgres, error) {
	c, err := pgxpool.ParseConfig(confString)
	if err != nil {
		return nil, err
	}

	return &Postgres{config: *c, ctx: ctx}, nil
}

func (p *Postgres) Connect() (*pgxpool.Pool, error) {
	pool, err := pgxpool.NewWithConfig(p.ctx, &p.config)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(p.ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
