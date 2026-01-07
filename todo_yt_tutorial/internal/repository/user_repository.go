package repository

import (
	// "context"
	// "time"
	"todo_api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUser(pool *pgxpool.Pool, user *models.User) (*models.User, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// query = "CREATE THE QUERY HERE"
	return &(models.User{}), nil
}
