package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"context"
	"log"
	"lumina/models"
)

func Connect() *pgxpool.Pool {
	db,err := pgxpool.New(context.Background(), "postgres://postgres:12345@localhost:5432/lumina")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetPosts(db *pgxpool.Pool) []models.ShowPost {
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
		show := models.ShowPost{
			Post: p,
			CreatedAtFormatted: p.CreatedAt.Format("01/02/2006"),
			UpdatedAtFormatted: p.UpdatedAt.Format("01/02/2006"),
		}
		posts = append(posts, show)
	}
	return posts
} 

func InsertUser(db *pgxpool.Pool, user models.User) {

}