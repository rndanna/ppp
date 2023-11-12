package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Postgres struct {
	Pool *sql.DB
}

func New(url string) (*Postgres, error) {
	pool, err := sql.Open("postgres",
		url)
	if err != nil {
		return nil, fmt.Errorf("Err Open: %w", err)
	}
	// defer pool.Close()

	return &Postgres{Pool: pool}, nil
}
