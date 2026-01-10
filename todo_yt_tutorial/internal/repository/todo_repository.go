package repository

import (
	"context"
	"errors"
	"time"
	"todo_api/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateTodo(pool *pgxpool.Pool, title string, completed bool, userid string) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var todo models.Todo
	query := `
		INSERT INTO todos_user (title, completed, user_id)
		VALUES ($1, $2, $3)
		RETURNING id, title, completed, created_at, updated_at, user_id
	`
	err := pool.QueryRow(ctx, query, title, completed, userid).Scan(
		&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.UserID,
	)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func GetAllTodos(pool *pgxpool.Pool, userid string) ([]models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "SELECT id, title, completed, created_at, updated_at, user_id FROM todos_user WHERE user_id = $1"
	rows, err := pool.Query(ctx, query, userid)
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
			&todo.UserID,
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

func GetTodoByID(pool *pgxpool.Pool, id int, userid string) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "SELECT id, title, completed, created_at, updated_at, user_id FROM todos_user WHERE id = $1 AND user_id = $2"
	var todo models.Todo
	err := pool.QueryRow(ctx, query, id, userid).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.UserID)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func UpdateTodoByID(pool *pgxpool.Pool, id int, title string, completed bool, userid string) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var todo models.Todo
	query := "UPDATE todos_user SET title=$1, completed=$2, updated_at=$3 WHERE id = $4 AND user_id = $5 RETURNING id, title, completed, created_at, updated_at, user_id"
	err := pool.QueryRow(ctx, query, title, completed, time.Now(), id, userid).Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.UserID)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func DeleteTodoByID(pool *pgxpool.Pool, id int, userid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "DELETE FROM todos_user WHERE id = $1 AND user_id = $2"
	commentTag, err := pool.Exec(ctx, query, id, userid)
	if err != nil {
		return err
	}

	if commentTag.RowsAffected() == 0 {
		return errors.New("todo not found")
	}
	return nil
}
