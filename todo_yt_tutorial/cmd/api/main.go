package main

import (
	"log"
	"todo_api/internal/config"
	"todo_api/internal/database"
	"todo_api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	var cfg *config.Config
	var err error
	cfg, err = config.Load()
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	pool, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	defer pool.Close()

	var router *gin.Engine = gin.Default()
	router.SetTrustedProxies(nil)
	router.POST("/todos", handlers.CreateTodoHandler(pool))
	router.GET("/todos/:id", handlers.GetTodoByIDHandler(pool))
	router.GET("/todos", handlers.GetAllTodosHandler(pool))
	router.PUT("/todos/:id", handlers.UpdateTodoByIDHandler(pool))
	router.DELETE("/todos/:id", handlers.DeleteTodoByIDHandler(pool))

	router.POST("/auth/register", handlers.CreateUserHandler(pool))
	router.Run(":300")
}
