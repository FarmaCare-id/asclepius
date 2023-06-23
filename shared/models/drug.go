package models

import (
	"gorm.io/gorm"
)


type (
	Drug struct {
		gorm.Model
		ID uint `gorm:"primaryKey;autoIncrement"`
		Code string `gorm:"column:code"`
		Name string `gorm:"column:name"`
		Description string `gorm:"column:description"`
	}
)