package models

import (
	"time"
)

type User struct {
	ID int
	CreatedAt time.Time
	UserID int64
	Stage string
}

type Student struct {
	ID int
	Name string
	ClassID int 
	Points string 
	CreatedAt time.Time
	TemplateID int
}

type Class struct {
	ID int 
	Name string 
	Grade int
	UserID int64
}

type Template struct {
	ID int
	Name string
	ClassID int
	UserID int64 
}