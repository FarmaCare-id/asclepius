package management

import (
	"farmacare/repository"
	"farmacare/shared"
	"farmacare/shared/dto"
	"farmacare/shared/models"
)

type (
	ViewService interface {
		CreateDrug(ctx dto.SessionContext, req dto.CreateDrugRequest) (dto.CreateDrugResponse, error)
	}

	viewService struct {
		repository  repository.Holder
		shared      shared.Holder
	}
)

func (v *viewService) CreateDrug(ctx dto.SessionContext, req dto.CreateDrugRequest) (dto.CreateDrugResponse, error) {
	var (
		res dto.CreateDrugResponse
	)

	drug := models.Drug {
		Code: req.Code,
		Name: req.Name,
		Description: req.Description,
	}

	err := v.repository.ManagementRepository.CreateDrug(drug)
	if err != nil {
		return res, err
	}

	res = dto.CreateDrugResponse{
		Code: req.Code,
		Name: req.Name,
		Description: req.Description,
	}

	return res, nil
}

func NewViewService(repository repository.Holder, shared shared.Holder) ViewService {
	return &viewService{
		repository: repository,
		shared:      shared,
	}
}