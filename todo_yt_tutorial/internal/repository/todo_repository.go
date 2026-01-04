package repository

import (
	"context"
	"time"
	"todo_api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTodo(pool *pgxpool.Pool, title string, completed bool) (models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var todo models.Todo
	query := `
		INSERT INTO todos_user (title, completed)
		VALUES ($1, $2)
		RETURNING id, title, completed, created_at, updated_at
	`
	err := pool.QueryRow(ctx, query, title, completed).Scan(
		&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt,
	)
	if err != nil {
		return models.Todo{}, err
	}
	return todo, nil
}
