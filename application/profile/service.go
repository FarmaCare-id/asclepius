package profile

import (
	"farmacare/shared"
	"farmacare/shared/models"
)

type (
	Service interface {
		EditUserProfile(user models.User) error
	}

	service struct {
		shared shared.Holder
	}
)

func (s *service) EditUserProfile(user models.User) error {
	err := s.shared.DB.Save(&user).Error
	return err
}

func NewProfileService(holder shared.Holder) (Service, error) {
	return &service{
		shared: holder,
	}, nil
}