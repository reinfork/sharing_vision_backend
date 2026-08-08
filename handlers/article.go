package handlers

import (
	"net/http"
	"strconv"

	"post-api/config"
	"post-api/models"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	if err := config.DB.Create(&post).Error; err != nil { // ≈ Model.create()
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": post})
}

func GetArticles(c *gin.Context) {
	limit, _ := strconv.Atoi(c.Param("limit"))
	offset, _ := strconv.Atoi(c.Param("offset"))
	if limit <= 0 {
		limit = 10
	}
	var posts []models.Post
	if err := config.DB.Limit(limit).Offset(offset).Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": posts})
}

func GetArticle(c *gin.Context) {
	var post models.Post
	if err := config.DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "article not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": post})
}

func UpdateArticle(c *gin.Context) {
	var post models.Post
	if err := config.DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "article not found"})
		return
	}
	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}
	if err := config.DB.Model(&post).Updates(payload).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": post})
}

func DeleteArticle(c *gin.Context) {
	var post models.Post
	if err := config.DB.First(&post, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "article not found"})
		return
	}
	if err := config.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "article deleted"})
}