package profile

import (
	"farmacare/shared"
	"farmacare/shared/dto"
)

type (
	Service interface {
		GetUserProfile(id uint) dto.User
	}

	service struct {
		shared shared.Holder
	}
)

func (s *service) GetUserProfile(id uint) dto.User {
	var user dto.User
	s.shared.DB.Where("id = ?", id).First(&user)
	return user
}

func NewProfileService(holder shared.Holder) (Service, error) {
	return &service{
		shared: holder,
	}, nil
}