package postgres

import "database/sql"

type Config struct {
	DSN string
}

func New(c Config) (*sql.DB, error) {
	return sql.Open("postgres", c.DSN)
}
