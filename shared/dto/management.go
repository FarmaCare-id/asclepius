package dto

import "farmacare/shared/models"

type (
	// GetAllDrugResponse GetAllDrugResponse
	GetAllDrugResponse struct {
		ID uint `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
		Description string `json:"description"`
	}

	// CreateDrugRequest CreateDrugRequest
	CreateDrugRequest struct {
		Code string `json:"code"`
		Name string `json:"name"`
		Description string `json:"description"`
	}

	// CreateDrugResponse CreateDrugResponse
	CreateDrugResponse struct {
		ID uint `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
		Description string `json:"description"`
	}
)

func (r *CreateDrugRequest) TransformToDrugModel() models.Drug {
	return models.Drug{
		Code: r.Code,
		Name: r.Name,
		Description: r.Description,
	}
}