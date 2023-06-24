package models

import (
	"gorm.io/gorm"
)


type (
	Delivery struct {
		gorm.Model
		ID uint `gorm:"primaryKey;autoIncrement"`
		DrugID uint
		Drug Drug `gorm:"onDelete:CASCADE"`
		UserID uint
		User User `gorm:"onDelete:CASCADE"`
		Status string `gorm:"column:status"`
		TrackingUrl string `gorm:"column:tracking_url"`
		Note string `gorm:"column:note"`
	}
)