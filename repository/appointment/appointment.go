package appointment

import (
	"farmacare/shared"
	"farmacare/shared/models"
)

type (
	Repository interface {
		CreateAppointment(appointment models.Appointment) error
		GetAllAppointmentByUserID(userId uint) ([]models.Appointment, error)
	}

	repository struct {
		shared shared.Holder
	}
)

func (s *repository) CreateAppointment(appointment models.Appointment) error {
	err := s.shared.DB.Create(&appointment).Error
	return err
}

func (s *repository) GetAllAppointmentByUserID(userId uint) ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := s.shared.DB.Where("patient_id = ?", userId).Find(&appointments).Error
	return appointments, err
}

func AppointmentRepository(holder shared.Holder) (Repository, error) {
	return &repository{
		shared: holder,
	}, nil
}