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
		CreatedAt time.Time
		ExpiredAt time.Time
	}
)