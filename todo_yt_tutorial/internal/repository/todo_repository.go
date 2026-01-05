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

func GetAllTodos(pool *pgxpool.Pool) ([]models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "SELECT id, title, completed, created_at, updated_at FROM todos_user"
	rows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}

func GetTodoByID(pool *pgxpool.Pool, id int) (models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "SELECT id, title, completed, created_at, updated_at FROM todos_user WHERE id = $1"
	var todo models.Todo
	err := pool.QueryRow(ctx, query, id).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return models.Todo{}, nil
		}
		return models.Todo{}, err
	}
	return todo, nil
}
