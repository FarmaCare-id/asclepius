package models

import (
	"gorm.io/gorm"
)

type (
	Forum struct {
		gorm.Model
		ID uint `gorm:"primaryKey;autoIncrement"`
		Title string `gorm:"column:title"`
		Description string `gorm:"column:description"`
		Vote uint `gorm:"column:vote"`
		UserID uint
		User User `gorm:"onDelete:CASCADE"`
		CreatedAt string `gorm:"column:created_at"`
		UpdatedAt string `gorm:"column:updated_at"`
		DeletedAt gorm.DeletedAt
	}
)