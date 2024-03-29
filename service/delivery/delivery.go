package delivery

import (
	"farmacare/repository"
	"farmacare/shared"
	"farmacare/shared/dto"
	"farmacare/shared/models"
)

type (
	ViewService interface {
		CreateDelivery(ctx dto.SessionContext, req dto.CreateDeliveryRequest) (dto.CreateDeliveryResponse, error)
		GetAllDeliveryForCurrentUser(ctx dto.SessionContext) ([]dto.GetDeliveryResponse, error)
	}

	viewService struct {
		repository  repository.Holder
		shared      shared.Holder
	}
)

func (v *viewService) CreateDelivery(ctx dto.SessionContext, req dto.CreateDeliveryRequest) (dto.CreateDeliveryResponse, error) {
	var (
		res dto.CreateDeliveryResponse
	)

	delivery := models.Delivery {
		DrugID: req.DrugID,
		UserID: ctx.User.ID,
		Status: req.Status,
		TrackingUrl: req.TrackingUrl,
		Note: req.Note,
	}

	err := v.repository.DeliveryRepository.CreateDelivery(delivery)
	if err != nil {
		return res, err
	}

	res = dto.CreateDeliveryResponse{
		ID: delivery.ID,
		DrugID: delivery.DrugID,
		UserID: delivery.UserID,
		Status: delivery.Status,
		TrackingUrl: delivery.TrackingUrl,
		Note: delivery.Note,
	}

	return res, nil
}

func (v *viewService) GetAllDeliveryForCurrentUser(ctx dto.SessionContext) ([]dto.GetDeliveryResponse, error) {
	var (
		res []dto.GetDeliveryResponse
	)

	deliveries, err := v.repository.DeliveryRepository.GetAllDeliveryByUserID(ctx.User.ID)
	if err != nil {
		return res, err
	}

	for _, delivery:= range deliveries {
		res = append(res, dto.GetDeliveryResponse{
			ID: delivery.ID,
			DrugID: delivery.DrugID,
			UserID: delivery.UserID,
			Status: delivery.Status,
			TrackingUrl: delivery.TrackingUrl,
			Note: delivery.Note,
		})
	}

	return res, nil
}

func NewViewService(repository repository.Holder, shared shared.Holder) ViewService {
	return &viewService{
		repository: repository,
		shared:      shared,
	}
}