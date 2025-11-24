package app

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"html/template"
)

var (
	DB *pgxpool.Pool  
	Tmpl *template.Template
)