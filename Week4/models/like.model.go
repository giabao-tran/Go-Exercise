package models

import (
	"gorm.io/gorm"
)

type Like struct {
	gorm.Model
	UserID uint `json:"user_id"`
	User   User `json:"user" gorm:"foreignKey:UserID"`
	PostID uint `json:"post_id"`
}