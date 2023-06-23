package management

import (
	"farmacare/shared"
	"farmacare/shared/models"
)

type (
	Repository interface {
		CheckDrugExist(code string) (bool, models.Drug)
		GetAllDrug(preload string) []models.Drug
		CreateDrug(drug models.Drug) error 
	}

	repository struct {
		shared shared.Holder
	}
)

func (s *repository) CheckDrugExist(code string) (bool, models.Drug) {
	var drug models.Drug
	err := s.shared.DB.First(&drug, "code = ?", code).Error
	return err == nil, drug
}

func (s *repository) GetAllDrug(preload string) []models.Drug {
	var drugs []models.Drug
	s.shared.DB.Preload(preload).Find(&drugs)
	return drugs
}

func (s *repository) CreateDrug(drug models.Drug) error {
	err := s.shared.DB.Create(&drug).Error
	return err
}

func ManagementRepository(holder shared.Holder) (Repository, error) {
	return &repository{
		shared: holder,
	}, nil
}