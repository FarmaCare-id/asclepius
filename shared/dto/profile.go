package dto

import "farmacare/shared/models"

type (
	// GetUserProfileRoleUserResponse GetUserProfileRoleUserResponse
	GetUserProfileRoleUserResponse struct {
		ID       uint `json:"id"`
		Email    string `json:"email"`
		Role     models.UserRole `json:"role"`
		Firstname string `json:"firstname"`
		Lastname string `json:"lastname"`
		Weight   float64 `json:"weight"`
		Height   float64 `json:"height"`
		Age      int `json:"age"`
	}

	// GetUserProfileRoleDoctorResponse GetUserProfileRoleDoctorResponse
	GetUserProfileRoleDoctorResponse struct {
		ID       uint `json:"id"`
		Email    string `json:"email"`
		Firstname string `json:"firstname"`
		Lastname string `json:"lastname"`
		Role     models.UserRole `json:"role"`
		NoSip    string `json:"no_sip"`
		Specialist string `json:"specialist"`
		Title    string `json:"title"`
	}

	// GetUserProfileRolePharmacistResponse GetUserProfileRolePharmacistResponse
	GetUserProfileRolePharmacistResponse struct {
		ID       uint `json:"id"`
		Email    string `json:"email"`
		Firstname string `json:"firstname"`
		Lastname string `json:"lastname"`
		Role     models.UserRole `json:"role"`
		NoSipa   string `json:"no_sipa"`
		Specialist string `json:"specialist"`
		Title    string `json:"title"`
	}
)