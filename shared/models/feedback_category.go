package models

import (
	"gorm.io/gorm"
)

type (
	FeedbackCategory struct {
		gorm.Model
		ID uint `gorm:"primaryKey;autoIncrement"`
		Name string `gorm:"column:name"`
		Description string `gorm:"column:description"`
		Feedbacks []Feedback `gorm:"foreignKey:FeedbackCategoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	}
)