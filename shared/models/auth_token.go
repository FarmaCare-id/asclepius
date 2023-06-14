package models

import (
	"time"
)

type (
	AuthToken struct {
		ID        uint `gorm:"primaryKey;autoIncrement"`
		UserID    uint
		User      User   `gorm:"onDelete:CASCADE"`
		Token     string `gorm:"column:token"`
		IsExpired bool   `gorm:"column:is_expired"`
		CreatedAt time.Time
	}
)