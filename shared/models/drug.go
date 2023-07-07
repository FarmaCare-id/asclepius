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
		Class string `gorm:"column:class"`
		Description string `gorm:"column:description"`
		Instruction string `gorm:"column:instruction"`
		Ingredient string `gorm:"column:ingredient"`
		Dose string `gorm:"column:dose"`
		Warning string `gorm:"column:warning"`	
	}
)