package models

import (
	"gorm.io/gorm"
)

type (
	ForumComment struct {
		gorm.Model
		ID uint `gorm:"primaryKey;autoIncrement"`
		UserID uint
		User User `gorm:"onDelete:CASCADE"`
		ForumID uint
		Forum Forum `gorm:"onDelete:CASCADE"`
		Comment string `gorm:"column:comment"`
		CreatedAt string `gorm:"column:created_at"`
		UpdatedAt string `gorm:"column:updated_at"`
		DeletedAt gorm.DeletedAt
	}
)