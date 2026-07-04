package database

import (
	"context"
	"github.com/Kashaan-Ekhlas/Key-Bored-Party/backend/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPool(cfg *config.Config) (*pgxpool.Pool, error) {

	poolConfig, err := pgxpool.ParseConfig(cfg.DatabaseURL())
	if err != nil {
		return nil, err
	}

	poolConfig.MaxConns = 20
	poolConfig.MinConns = 5

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		pool.Close()
		return nil, err
	}
	return pool, nil
}
