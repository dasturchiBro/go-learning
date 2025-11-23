package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
	"log"
	"lumina/models"
)

func connect() (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), "postgres://postgres:12345@localhost:5432/lumina")
}

func insertUser(db *pgxpool.Pool, user models.User) {

}