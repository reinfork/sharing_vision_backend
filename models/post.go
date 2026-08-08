package models

import "time"

// Post mirrors your MySQL table 1:1 via gorm tags.
type Post struct {
	ID          int       `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	Title       string    `gorm:"column:title;size:200" json:"title" binding:"required,min=20"`
	Content     string    `gorm:"column:content;type:text" json:"content" binding:"required,min=200"`
	Category    string    `gorm:"column:category;size:100" json:"category" binding:"required,min=3"`
	CreatedDate time.Time `gorm:"column:created_date;autoCreateTime" json:"created_date"`
	UpdatedDate time.Time `gorm:"column:updated_date;autoUpdateTime" json:"updated_date"`
	Status      string    `gorm:"column:status;size:100" json:"status" binding:"required,oneof=Publish Draft Thrash"`
}
