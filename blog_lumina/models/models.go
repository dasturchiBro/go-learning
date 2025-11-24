package models

import (
	"time"
)

type User struct {
	Id int 
	Name string
	Email string
	Hash_pass string
	Role string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Post struct {
	Id int
	Title string
	Body string
	UserId int
	Category string
	Image string
	Views int  
	CreatedAt time.Time
	UpdatedAt time.Time
	IsPublished bool 
}

type ShowPost struct {
	Post Post
	CreatedAtFormatted string
	UpdatedAtFormatted string
}