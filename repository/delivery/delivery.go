package delivery

import (
	"farmacare/shared"
	"farmacare/shared/models"
)

type (
	DeliveryRepository interface {
		CreateDelivery(delivery models.Delivery) error 
	}

	deliveryRepository struct {
		shared shared.Holder
	}
)

func (s *deliveryRepository) CreateDelivery(delivery models.Delivery) error {
	err := s.shared.DB.Create(&delivery).Error
	return err
}

func NewDeliveryRepository(holder shared.Holder) (DeliveryRepository, error) {
	return &deliveryRepository{
		shared: holder,
	}, nil
}