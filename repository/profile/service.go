package profile

import (
	"farmacare/shared"
	"farmacare/shared/models"
)

type (
	Repository interface {
		EditUserProfile(user models.User) error
	}

	repository struct {
		shared shared.Holder
	}
)

func (s *repository) EditUserProfile(user models.User) error {
	err := s.shared.DB.Save(&user).Error
	return err
}

func ProfileRepository(holder shared.Holder) (Repository, error) {
	return &repository{
		shared: holder,
	}, nil
}