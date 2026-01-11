package main

import (
	"log"
	"todo_api/internal/config"
	"todo_api/internal/database"
	"todo_api/internal/handlers"
	"todo_api/internal/middleware"
	"time"

	"github.com/gin-contrib/cors"
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
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	router.Use(middleware.CorsMiddleware())

	router.OPTIONS("/", func(c *gin.Context) {
		c.Status(204)
	})
	router.POST("/auth/register", handlers.CreateUserHandler(pool))
	router.POST("/auth/login", handlers.LoginUserHandler(pool, cfg))

	protected := router.Group("/todos")

	protected.Use(middleware.AuthMiddleware(cfg))

	protected.OPTIONS("", func(c *gin.Context) {
		c.Status(204)
	})
	protected.OPTIONS("/:id", func(c *gin.Context) {
		c.Status(204)
	})
	protected.POST("", handlers.CreateTodoHandler(pool))
	protected.GET("/:id", handlers.GetTodoByIDHandler(pool))
	protected.GET("", handlers.GetAllTodosHandler(pool))
	protected.PUT("/:id", handlers.UpdateTodoByIDHandler(pool))
	protected.DELETE("/:id", handlers.DeleteTodoByIDHandler(pool))

	// Middleware Test Route
	router.GET("/test", handlers.TestProtectedHandler())

	router.Run(":3000")
}
