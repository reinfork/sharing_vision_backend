package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	router = gin.Default()

	router.POST("/article/", CreateArticle)
	router.GET("/article/:id", GetArticle)
	router.GET("/article/:id/:offset", GetArticles)
	router.PUT("/article/:id", UpdateArticle)
	router.DELETE("/article/:id", DeleteArticle)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// This passes the incoming Vercel HTTP request directly into your Gin router
	router.ServeHTTP(w, r)
}