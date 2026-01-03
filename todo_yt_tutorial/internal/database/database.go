package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(databaseURL string) (*pgxpool.Pool, error) {
	ctx := context.Background()

	config, err := pgxpool.ParseConfig(databaseURL)

	if err != nil {
		// Log the error and return
		println("Failed to parse database URL:", err.Error())
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		println("Failed to create connection pool:", err.Error())
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		println("Failed to ping database:", err.Error())
		pool.Close()
		return nil, err
	}

	log.Println("Successfully connected to the database")
	return pool, nil
}
