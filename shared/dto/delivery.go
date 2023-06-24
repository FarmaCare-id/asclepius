package dto

type (
	// CreateDeliveryRequest CreateDeliveryRequest
	CreateDeliveryRequest struct {
		DrugID uint `json:"drug_id"`
		Status string `json:"status"`
		TrackingUrl string `json:"tracking_url"`
		Note string `json:"note"`
	}
		
	// CreateDeliveryResponse CreateDeliveryResponse
	CreateDeliveryResponse struct {
		ID uint `json:"id"`
		DrugID uint `json:"drug_id"`
		UserID uint `json:"user_id"`
		Status string `json:"status"`
		TrackingUrl string `json:"tracking_url"`
		Note string `json:"note"`
	}

	// GetDeliveryResponse GetDeliveryResponse
	GetDeliveryResponse struct {
		ID uint `json:"id"`
		DrugID uint `json:"drug_id"`
		UserID uint `json:"user_id"`
		Status string `json:"status"`
		TrackingUrl string `json:"tracking_url"`
		Note string `json:"note"`
	}
)