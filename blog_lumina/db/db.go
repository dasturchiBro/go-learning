package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
	"log"
	"lumina/models"
	"lumina/app"
	"time"
)

func Connect() *pgxpool.Pool {
	db,err := pgxpool.New(context.Background(), "postgres://postgres:12345@localhost:5432/lumina")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetPosts() []models.ShowPost {
	db := app.DB
	rows, err := db.Query(
		context.Background(), 
		"SELECT id, title, body, user_id, category, image, views, created_at, updated_at, is_published FROM posts",
	)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	var posts []models.ShowPost
	for rows.Next() {
		var p models.Post
		err := rows.Scan(&p.Id,&p.Title, &p.Body, &p.UserId, &p.Category, &p.Image, &p.Views, &p.CreatedAt, &p.UpdatedAt, &p.IsPublished)
		if err != nil {
			log.Fatal(err)
		}
		row := db.QueryRow(context.Background(), "SELECT id, name FROM users WHERE id = $1", p.UserId)
		var newUser models.User
		err = row.Scan(&newUser.Id, &newUser.Name)
		if err != nil {
			log.Fatal(err)
		}
		show := models.ShowPost{
			Post: p,
			CreatedAtFormatted: p.CreatedAt.Format("01/02/2006"),
			UpdatedAtFormatted: p.UpdatedAt.Format("01/02/2006"),
			User: newUser,
		}
		posts = append(posts, show)
	}
	return posts
} 

func GetPostById(id int) (models.ShowPost, bool) {
	rows, err := app.DB.Query(context.Background(), "SELECT * FROM posts WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var show models.ShowPost
	if !rows.Next() {
		return show, false
	}
	var post models.Post
	err = rows.Scan(&post.Id,&post.Title, &post.Body, &post.UserId, &post.Category, &post.Image, &post.Views, &post.CreatedAt, &post.UpdatedAt, &post.IsPublished)
	if err != nil {
		log.Fatal(err)
	}
	row := app.DB.QueryRow(context.Background(), "SELECT id, name FROM users WHERE id = $1", post.UserId)
	var newUser models.User
	err = row.Scan(&newUser.Id, &newUser.Name)
	log.Print(newUser)
	if err != nil {
		log.Fatal(err)
	}
	show.Post = post
	show.CreatedAtFormatted = post.CreatedAt.Format("01/02/2006")
	show.UpdatedAtFormatted = post.UpdatedAt.Format("01/02/2006")
	show.User = newUser
	return show, true
}



func InsertUser(user models.User) (int, string) {
	db := app.DB  
	rows, err := db.Query(context.Background(), "SELECT id FROM users WHERE email = $1", user.Email)
	if err != nil {
		return 0, "error"
	}
	if rows.Next() {
		return 0, "email"
	}
	row := db.QueryRow(context.Background(), "INSERT INTO users (name, email, hash_pass, role, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", user.Name, user.Email, user.Hash_pass, user.Role, user.CreatedAt, user.UpdatedAt)
	var UserId int
	err = row.Scan(&UserId)
	if err != nil {
		log.Print(UserId, err)
		return 0, "error"
	}
	return UserId, "success"
}

func GetSessionExpiryTime(session models.Session) (time.Time, error) {
	db := app.DB 
	var expiryTime time.Time
	row := db.QueryRow(context.Background(), "SELECT session_id, user_id, expires_at FROM sessions WHERE session_id = $1 AND user_id = $2", session.SessionId, session.UserId)
	err := row.Scan(&expiryTime)
	return expiryTime, err
}

func InsertSession(session models.Session) error {
	_, err := app.DB.Exec(context.Background(), "INSERT INTO sessions (session_id, user_id, expires_at) VALUES ($1, $2, $3)", session.SessionId, session.UserId, session.ExpiresAt)
	return err
}

func DeleteSession(session models.Session) error {
	db := app.DB  
	_, err := db.Exec(context.Background(), "DELETE FROM sessions WHERE session_id = $1 AND user_id = $2", session.SessionId, session.UserId)
	return err
}