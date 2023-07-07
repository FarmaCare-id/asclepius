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
		Class string `json:"class"`
		Description string `json:"description"`
		Instruction string `json:"instruction"`
		Ingredient string `json:"ingredient"`
		Dose string `json:"dose"`
		Warning string `json:"warning"`
	}

	// CreateDrugResponse CreateDrugResponse
	CreateDrugResponse struct {
		ID uint `json:"id"`
		Code string `json:"code"`
		Name string `json:"name"`
		Class string `json:"class"`
		Description string `json:"description"`
		Instruction string `json:"instruction"`
		Ingredient string `json:"ingredient"`
		Dose string `json:"dose"`
		Warning string `json:"warning"`
	}

	// CreateUserDrugRequest CreateUserDrugRequest
	CreateUserDrugRequest struct {
		UserID uint `json:"user_id"`
		DrugID uint `json:"drug_id"`
		Status string `json:"status"`
		Note string `json:"note"`
		FrequencyPerDay uint `json:"frequency_per_day"`
	}

	// CreateUserDrugResponse CreateUserDrugResponse
	CreateUserDrugResponse struct {
		ID uint `json:"id"`
		UserID uint `json:"user_id"`
		DrugID uint `json:"drug_id"`
		Status string `json:"status"`
		Note string `json:"note"`
		FrequencyPerDay uint `json:"frequency_per_day"`
	}

	// GetUserDrugResponse GetUserDrugResponse
	GetUserDrugResponse struct {
		ID uint `json:"id"`
		UserID uint `json:"user_id"`
		DrugID uint `json:"drug_id"`
		Status string `json:"status"`
		Note string `json:"note"`
		FrequencyPerDay uint `json:"frequency_per_day"`
	}
)

func (r *CreateDrugRequest) TransformToDrugModel() models.Drug {
	return models.Drug{
		Code: r.Code,
		Name: r.Name,
		Class: r.Class,
		Description: r.Description,
		Instruction: r.Instruction,
		Ingredient: r.Ingredient,
		Dose: r.Dose,
		Warning: r.Warning,
	}
}