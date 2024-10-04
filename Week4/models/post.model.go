package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	UserID   uint      `json:"user_id"`
	User     User      `json:"user" gorm:"foreignKey:UserID"`
	Comments []Comment `json:"comments" gorm:"foreignKey:PostID"`
	Likes    []Like    `json:"likes" gorm:"foreignKey:PostID"`
}
