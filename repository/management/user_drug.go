package management

import (
	"farmacare/shared"
	"farmacare/shared/models"
)

type (
	UserDrugRepository interface {
		CreateUserDrug(userDrug models.UserDrug) error
	}

	userDrugRepository struct {
		shared shared.Holder
	}
)

func (s *userDrugRepository) CreateUserDrug(userDrug models.UserDrug) error {
	err := s.shared.DB.Create(&userDrug).Error
	return err
}

func NewUserDrugRepository(holder shared.Holder) (UserDrugRepository, error) {
	return &userDrugRepository{
		shared: holder,
	}, nil
}