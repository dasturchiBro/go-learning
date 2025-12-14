package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
	"os"

	"xlsxbot/models"
	"xlsxbot/app"
)

func Connect() *pgxpool.Pool {
	return pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
}

func SelectUserByUserID(user_id string) (int, error) {
	var id int
	err := app.DB.QueryRow(context.Background(), "SELECT id FROM users WHERE user_id = $1", user_id).Scan(&id) 
	if err != nil {
		return nil, err
	}
	return id, nil
}