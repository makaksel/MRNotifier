package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	DSN string
}

func New(c Config) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), c.DSN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	return pool
}
