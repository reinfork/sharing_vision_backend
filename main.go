package main

import (
	"post-api/config"
	"post-api/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	config.ConnectDB()
	r := gin.Default()
	setupRoutes(r)
	r.Run(":8080")
}

func setupRoutes(r *gin.Engine) {
	r.Use(corsMiddleware())

	article := r.Group("/article")
	{
		article.POST("/", handlers.CreateArticle)
		article.GET("/:limit/:offset", handlers.GetArticles)
		article.GET("/:id", handlers.GetArticle)
		article.PUT("/:id", handlers.UpdateArticle)
		article.DELETE("/:id", handlers.DeleteArticle)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}