package models

import (
	"gorm.io/gorm"
)

type (
	UserDrug struct {
		gorm.Model
		ID uint `gorm:"primaryKey;autoIncrement"`
		UserID uint
		User User `gorm:"onDelete:CASCADE"`
		DrugID uint
		Drug Drug `gorm:"onDelete:CASCADE"`
		Status string `gorm:"column:status"`
		Note string `gorm:"column:note"`
		FrequencyPerDay uint `gorm:"column:frequency"`
	}
)