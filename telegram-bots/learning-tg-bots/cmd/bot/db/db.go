package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
	"os"
	"xlsxbot/models"
	"xlsxbot/app"
	"errors"
	"strconv"
)

func Connect() (*pgxpool.Pool, error) {
	return pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
}

func UserExistsByUserID(user_id int64) (bool, error) {
	var exists bool
	err := app.DB.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1)", user_id).Scan(&exists) 
	return exists, err
}

func InsertUser(user_id int64) (int, error) {
	var id int 
	query := `INSERT INTO users (user_id, stage) VALUES ($1, $2) RETURNING id`
	err := app.DB.QueryRow(context.Background(), query, user_id, "main").Scan(&id)
	return id, err
}

func GetStageByUserID(chatID int64) (string, error) {
	query := "SELECT stage FROM users where user_id = $1"
	var stage string
	err := app.DB.QueryRow(context.Background(), query, chatID).Scan(&stage)
	return stage, err
}

func SetStageByUserID(chatID int64, value string) (bool, error) {
	exists, err := UserExistsByUserID(chatID)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, nil
	}
	query := "UPDATE users SET stage = $1 WHERE user_id = $2 RETURNING stage"
	var stage string
	err = app.DB.QueryRow(context.Background(), query, value, chatID).Scan(&stage)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetClassesByUserID(chatID int64) ([]models.Class, error) {
	query := "SELECT * FROM classes WHERE user_id = $1"
	rows, err := app.DB.Query(context.Background(), query, chatID)
	if err != nil {
		return nil, err
	}
	classes := make([]models.Class, 0)

	for rows.Next() {
		var class models.Class
		err := rows.Scan(&class.ID, &class.Name, &class.Grade, &class.UserID)
		if err != nil {
			return nil, err
		}
		classes = append(classes, class)

	}
	return classes, nil
}

func AddClass(chatID int64, className string, grade int) (int, error) {
	query := "INSERT INTO classes (name, grade, user_id) VALUES ($1, $2, $3) RETURNING id"
	var id int
	err := app.DB.QueryRow(context.Background(), query, className, grade, chatID).Scan(&id)
	return id, err
}

func RemoveClass(chatID int64, classID int) (bool, error) {
	query := "DELETE FROM classes WHERE user_id = $1 AND id = $2"
	exec, err := app.DB.Exec(context.Background(), query, chatID, classID)
	if err != nil {
		return false, err
	}
	if exec.RowsAffected() != 0 {
		return true, nil
	}
	return false, nil
}

func GetTemplatesByUserID(chatID int64) ([]models.Template, error) {
	query := "SELECT * FROM templates WHERE user_id = $1"
	rows, err := app.DB.Query(context.Background(), query, chatID)
	if err != nil {
		return nil, err
	}
	templates := make([]models.Template, 0)

	for rows.Next() {
		var template models.Template
		err := rows.Scan(&template.ID, &template.Name, &template.ClassID, &template.UserID)
		if err != nil {
			return nil, err
		}
		templates = append(templates, template)

	}
	return templates, nil
}

func GetClassByUserID(chatID int64, id int) (models.Class, error) {
	query := "SELECT * FROM classes WHERE user_id = $1 AND id = $2"
	var class models.Class
	rows, err := app.DB.Query(context.Background(), query, chatID, id)
	if err != nil {
		return class, err
	}
	if rows.Next() {
		err := rows.Scan(&class.ID, &class.Name, &class.Grade, &class.UserID)
		if err != nil {
			return class, err
		}
	} else {
		return class, errors.New("Class with ID" + strconv.Itoa(id) + "doesn't exist")
	}
	return class, nil
}