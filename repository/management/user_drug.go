package management

import (
	"farmacare/shared"
	"farmacare/shared/models"
)

type (
	UserDrugRepository interface {
		CreateUserDrug(userDrug models.UserDrug) error
		GetAllUserDrugByUserID(userId uint) ([]models.UserDrug, error)
	}

	userDrugRepository struct {
		shared shared.Holder
	}
)

func (s *userDrugRepository) CreateUserDrug(userDrug models.UserDrug) error {
	err := s.shared.DB.Create(&userDrug).Error
	return err
}

func (s *userDrugRepository) GetAllUserDrugByUserID(userId uint) ([]models.UserDrug, error) {
	var userDrugs []models.UserDrug
	err := s.shared.DB.Where("user_id = ?", userId).Find(&userDrugs).Error
	return userDrugs, err
}

func NewUserDrugRepository(holder shared.Holder) (UserDrugRepository, error) {
	return &userDrugRepository{
		shared: holder,
	}, nil
}