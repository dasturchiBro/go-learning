package models

import (
	"time"
)

type User struct {
	Id int
	UserID string
	Stage string
	CreatedAt time.Time
}

type Student struct {
	Id int
	Name string
	ClassID int 
	Points string 
	CreatedAt time.Time
}

type Class struct {
	Id int 
	Name string 
	Grade int
}