package models

import (
	"gorm.io/gorm"
)

type (
	Appointment struct {
		gorm.Model
		ID uint `gorm:"primaryKey;autoIncrement"`
		PatientID uint
		Patient User `gorm:"onDelete:CASCADE"`
		HealthcareWorkerID uint
		HealthcareWorker User `gorm:"onDelete:CASCADE"`
		Complaint string `gorm:"column:complaint"`
		ScheduleDate string `gorm:"column:schedule_date"`
		ScheduleTime string `gorm:"column:schedule_time"`
		Location string `gorm:"column:location"`
		Status string `gorm:"column:status"`
		Note string `gorm:"column:note"`
	}
)