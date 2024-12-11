package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresPool struct {
	Pool *pgxpool.Pool
}

func NewPostgresPool(dsn string) *PostgresPool {
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	fmt.Println("Connected to PostgreSQL! 🐘")
	return &PostgresPool{Pool: pool}
}

func (p *PostgresPool) Close() {
	p.Pool.Close()
}
