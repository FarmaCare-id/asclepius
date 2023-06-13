package models

import (
	"time"
)

type (
	PasswordReset struct {
		ID        uint `gorm:"primaryKey;autoIncrement"`
		UserID    uint
		User      User   `gorm:"onDelete:CASCADE"`
		Token     string `gorm:"column:token"`
		CreatedAt time.Time
		Valid     time.Time
	}	
)