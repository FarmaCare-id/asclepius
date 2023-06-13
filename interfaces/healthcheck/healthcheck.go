package healthcheck

import (
	"farmacare/repository"
	"farmacare/shared"
	"farmacare/shared/dto"
)

type (
	ViewService interface {
		SystemHealthcheck() (dto.HCStatus, error)
	}

	viewService struct {
		repository repository.Holder
		shared      shared.Holder
	}
)

func (v *viewService) SystemHealthcheck() (dto.HCStatus, error) {
	status := make([]dto.Status, 0)

	httpStatus := v.repository.HealthcheckService.HttpHealthcheck(v.shared.Http)
	status = append(status, httpStatus)

	dbStatus := v.repository.HealthcheckService.DatabaseHealthcheck(v.shared.DB)
	status = append(status, dbStatus)

	return dto.HCStatus{
		Status: status,
	}, nil
}

func NewViewService(repository repository.Holder, shared shared.Holder) ViewService {
	return &viewService{
		repository: repository,
		shared:      shared,
	}
}