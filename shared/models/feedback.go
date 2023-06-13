package models

import (
	"time"

	"gorm.io/gorm"
)


type (
	Feedback struct {
		gorm.Model
		ID uint `gorm:"primaryKey;autoIncrement"`
		Issue string `gorm:"column:issue"`
		Detail string `gorm:"column:detail"`
		FeedbackCategoryID uint
		FeedbackCategory FeedbackCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		UserID uint
		User User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt
	}
)