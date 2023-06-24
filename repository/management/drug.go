package management

import (
	"farmacare/shared"
	"farmacare/shared/models"
)

type (
	DrugRepository interface {
		CheckDrugExist(code string) (bool, models.Drug)
		GetAllDrug(preload string) []models.Drug
		CreateDrug(drug models.Drug) error 
		GetDrugByID(id uint) (models.Drug, error)
	}

	drugRepository struct {
		shared shared.Holder
	}
)

func (s *drugRepository) CheckDrugExist(code string) (bool, models.Drug) {
	var drug models.Drug
	err := s.shared.DB.First(&drug, "code = ?", code).Error
	return err == nil, drug
}

func (s *drugRepository) GetAllDrug(preload string) []models.Drug {
	var drugs []models.Drug
	s.shared.DB.Preload(preload).Find(&drugs)
	return drugs
}

func (s *drugRepository) CreateDrug(drug models.Drug) error {
	err := s.shared.DB.Create(&drug).Error
	return err
}

func (s *drugRepository) GetDrugByID(id uint) (models.Drug, error) {
	var drug models.Drug
	err := s.shared.DB.First(&drug, id).Error
	return drug, err
}

func NewDrugRepository(holder shared.Holder) (DrugRepository, error) {
	return &drugRepository{
		shared: holder,
	}, nil
}