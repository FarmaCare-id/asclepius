package models

import (
	"time"

	"gorm.io/gorm"
)

const (
	Google     UserLoginType = "Google"
	Credential UserLoginType = "Credential"
)

const (
	RoleUser       UserRole = "user"
	RoleDoctor     UserRole = "doctor"
	RolePharmacist UserRole = "pharmacist"
	RoleAdmin      UserRole = "admin"
)

type (
	UserLoginType string
	UserRole string
	User          struct {
		ID                uint          `gorm:"primaryKey;autoIncrement"`
		Email             string        `gorm:"column:email;unique;not null"`
		Fullname          string        `gorm:"column:fullname"`
		HashedPassword    string        `gorm:"column:hashed_password"`
		Role           	  UserRole      `gorm:"column:role;default:user"`
		Type              UserLoginType `gorm:"column:type;default:Credential"`
		Weight			  float64		`gorm:"column:weight"`
		Height			  float64		`gorm:"column:height"`
		Age				  int			`gorm:"column:age"`
		NoSip			  string		`gorm:"column:no_sip"`
		NoSipa			  string		`gorm:"column:no_sipa"`
		Specialist		  string		`gorm:"column:specialist"`
		Title 			  string		`gorm:"column:title"`
		CreatedAt         time.Time
		UpdatedAt         time.Time
		DeletedAt         gorm.DeletedAt
	}
)