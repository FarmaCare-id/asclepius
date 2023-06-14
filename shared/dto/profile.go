package dto

import "farmacare/shared/models"

type (
	// GetUserProfileRoleUserResponse GetUserProfileRoleUserResponse
	GetUserProfileRoleUserResponse struct {
		ID       uint `json:"id"`
		Email    string `json:"email"`
		Role     models.UserRole `json:"role"`
		Fullname string `json:"fullname"`
		Weight   float64 `json:"weight"`
		Height   float64 `json:"height"`
		Age      int `json:"age"`
	}

	// GetUserProfileRoleDoctorResponse GetUserProfileRoleDoctorResponse
	GetUserProfileRoleDoctorResponse struct {
		ID       uint `json:"id"`
		Email    string `json:"email"`
		Fullname string `json:"fullname"`
		Role     models.UserRole `json:"role"`
		NoSip    string `json:"no_sip"`
		Specialist string `json:"specialist"`
		Title    string `json:"title"`
	}

	// GetUserProfileRolePharmacistResponse GetUserProfileRolePharmacistResponse
	GetUserProfileRolePharmacistResponse struct {
		ID       uint `json:"id"`
		Email    string `json:"email"`
		Fullname string `json:"fullname"`
		Role     models.UserRole `json:"role"`
		NoSipa   string `json:"no_sipa"`
		Specialist string `json:"specialist"`
		Title    string `json:"title"`
	}

	// GetUserProfileRoleAdminResponse GetUserProfileRoleAdminResponse
	GetUserProfileRoleAdminResponse struct {
		ID       uint `json:"id"`
		Email    string `json:"email"`
		Fullname string `json:"fullname"`
		Role     models.UserRole `json:"role"`
		Weight   float64 `json:"weight"`
		Height   float64 `json:"height"`
		Age      int `json:"age"`
		NoSip    string `json:"no_sip"`
		NoSipa   string `json:"no_sipa"`
		Specialist string `json:"specialist"`
		Title    string `json:"title"`
	}
)