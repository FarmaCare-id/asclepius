package management

import (
	"farmacare/repository"
	"farmacare/shared"
	"farmacare/shared/dto"
	"farmacare/shared/exception"
	"farmacare/shared/models"
)

type (
	ViewService interface {
		GetAllDrug() ([]dto.GetAllDrugResponse, error)
		CreateDrug(ctx dto.SessionContext, req dto.CreateDrugRequest) (dto.CreateDrugResponse, error)
		CreateUserDrug(ctx dto.SessionContext, req dto.CreateUserDrugRequest) (dto.CreateUserDrugResponse, error)
	}

	viewService struct {
		repository  repository.Holder
		shared      shared.Holder
	}
)

func (v *viewService) GetAllDrug() ([]dto.GetAllDrugResponse, error) {
	var (
		res []dto.GetAllDrugResponse
	)

	drugs := v.repository.DrugRepository.GetAllDrug("")
	for _, drug:= range drugs {
		res = append(res, dto.GetAllDrugResponse{
			ID: drug.ID,
			Code: drug.Code,
			Name: drug.Name,
		})
	}

	return res, nil
}



func (v *viewService) CreateDrug(ctx dto.SessionContext, req dto.CreateDrugRequest) (dto.CreateDrugResponse, error) {
	var (
		res dto.CreateDrugResponse
	)

	isDrugExist, _ := v.repository.DrugRepository.CheckDrugExist(req.Code)
	if isDrugExist {
		return res, exception.DrugAlreadyExist()
	}

	drug := models.Drug {
		Code: req.Code,
		Name: req.Name,
		Description: req.Description,
	}

	err := v.repository.DrugRepository.CreateDrug(drug)
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

func (v *viewService) CreateUserDrug(ctx dto.SessionContext, req dto.CreateUserDrugRequest) (dto.CreateUserDrugResponse, error) {
	var (
		res dto.CreateUserDrugResponse
	)

	drug, err := v.repository.DrugRepository.GetDrugByID(req.DrugID)
	if err != nil {
		return res, exception.DrugNotFound()
	}

	userDrug := models.UserDrug {
		UserID: ctx.User.ID,
		DrugID: drug.ID,
		Status: req.Status,
		Note: req.Note,
		FrequencyPerDay: req.FrequencyPerDay,
	}

	err = v.repository.UserDrugRepository.CreateUserDrug(userDrug)
	if err != nil {
		return res, err
	}

	res = dto.CreateUserDrugResponse{
		UserID: ctx.User.ID,
		DrugID: drug.ID,
		Status: req.Status,
		Note: req.Note,
		FrequencyPerDay: req.FrequencyPerDay,
	}

	return res, nil
}


func NewViewService(repository repository.Holder, shared shared.Holder) ViewService {
	return &viewService{
		repository: repository,
		shared:      shared,
	}
}