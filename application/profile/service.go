package profile

import (
	"farmacare/shared"
	"farmacare/shared/dto"
)

type (
	Service interface {
		EditUserProfile(user dto.User) error
	}

	service struct {
		shared shared.Holder
	}
)

func (s *service) EditUserProfile(user dto.User) error {
	err := s.shared.DB.Save(&user).Error
	return err
}

func NewProfileService(holder shared.Holder) (Service, error) {
	return &service{
		shared: holder,
	}, nil
}