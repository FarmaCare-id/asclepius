package delivery

import (
	"farmacare/shared"
	"farmacare/shared/models"
)

type (
	DeliveryRepository interface {
		CreateDelivery(delivery models.Delivery) error 
		GetAllDeliveryByUserID(userID uint) ([]models.Delivery, error)
		GetAllDeliveryForCurrentUser(preload string) []models.Delivery
	}

	deliveryRepository struct {
		shared shared.Holder
	}
)

func (s *deliveryRepository) CreateDelivery(delivery models.Delivery) error {
	err := s.shared.DB.Create(&delivery).Error
	return err
}

func (s *deliveryRepository) GetAllDeliveryByUserID(userID uint) ([]models.Delivery, error) {
	var deliveries []models.Delivery
	err := s.shared.DB.Where("user_id = ?", userID).Find(&deliveries).Error
	return deliveries, err
}

func (s *deliveryRepository) GetAllDeliveryForCurrentUser(preload string) []models.Delivery {
	var deliveries []models.Delivery
	s.shared.DB.Preload(preload).Find(&deliveries)
	return deliveries
}

func NewDeliveryRepository(holder shared.Holder) (DeliveryRepository, error) {
	return &deliveryRepository{
		shared: holder,
	}, nil
}