package dto

import (
	"farmacare/shared/models"
)

type (
	// CreateUserRequest CreateUserRequest
	CreateUserRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Firstname string `json:"firstname" validate:"required"`
		Lastname string `json:"lastname" validate:"required"`
	}

	// CreateUserResponse CreateUserResponse
	CreateUserResponse struct {
		Email    string `json:"email"`
		Firstname string `json:"firstname"`
		Lastname string `json:"lastname"`
		Role     string `json:"role"`
	}

	// CreateDoctorRequest CreateDoctorRequest
	CreateDoctorRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Firstname string `json:"firstname" validate:"required"`
		Lastname string `json:"lastname" validate:"required"`
		NoSip    string `json:"no_sip" validate:"required"`
	}

	// CreateDoctorResponse CreateDoctorResponse
	CreateDoctorResponse struct {
		Email    string `json:"email"`
		Firstname string `json:"firstname"`
		Lastname string `json:"lastname"`
		Role     string `json:"role"`
		NoSip    string `json:"no_sip"`
	}

	// CreatePharmacistRequest CreatePharmacistRequest
	CreatePharmacistRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Firstname string `json:"firstname" validate:"required"`
		Lastname string `json:"lastname" validate:"required"`
		NoSipa   string `json:"no_sipa" validate:"required"`
	}

	// CreatePharmacistResponse CreatePharmacistResponse
	CreatePharmacistResponse struct {
		Email    string `json:"email"`
		Firstname string `json:"firstname"`
		Lastname string `json:"lastname"`
		Role     string `json:"role"`
		NoSipa   string `json:"no_sipa"`
	}

	// CreateAdminRequest CreateAdminRequest
	CreateAdminRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Firstname string `json:"firstname" validate:"required"`
		Lastname string `json:"lastname" validate:"required"`
	}

	// CreateAdminResponse CreateAdminResponse
	CreateAdminResponse struct {
		Email    string `json:"email"`
		Firstname string `json:"firstname"`
		Lastname string `json:"lastname"`
		Role     string `json:"role"`
	}
		
	// EditUserRequest EditUserRequest
	EditUserRequest struct {
		Firstname string `json:"firstname" validate:"omitempty"`
		Lastname string `json:"lastname" validate:"omitempty"`
		Password string `json:"password" validate:"omitempty"`
		Weight   float64 `json:"weight" validate:"omitempty"`
		Height   float64 `json:"height" validate:"omitempty"`
		Age      int `json:"age" validate:"omitempty"`
		NoSip    string `json:"no_sip" validate:"omitempty"`
		NoSipa   string `json:"no_sipa" validate:"omitempty"`
		Specialist string `json:"specialist" validate:"omitempty"`
		Title    string `json:"title" validate:"omitempty"`
	}

	// EditUserResponse EditUserResponse
	EditUserResponse struct {
		Firstname string `json:"firstname"`
		Lastname string `json:"lastname"`
		Weight   float64 `json:"weight"`
		Height   float64 `json:"height"`
		Age      int `json:"age"`
		NoSip    string `json:"no_sip"`
		NoSipa   string `json:"no_sipa"`
		Specialist string `json:"specialist"`
		Title    string `json:"title"`
	}

	// LoginRequest LoginRequest
	LoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	// LoginResponse LoginResponse
	LoginResponse struct {
		Token string `json:"token"`
	}

	// GoogleLoginRequest GoogleLoginRequest
	GoogleLoginRequest struct {
		Token string `json:"token" validate:"required"`
	}

	GoogleData struct {
		Sub           string `json:"sub"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Picture       string `json:"picture"`
		Email         string `json:"email"`
		EmailVerified string `json:"email_verified"`
		Locale        string `json:"locale"`
	}

	// LogoutResponse LogoutResponse
	LogoutResponse struct {
		Token     string `json:"token"`
	}

	// ForgotPasswordRequest ForgotPasswordRequest
	ForgotPasswordRequest struct {
		Email string `json:"email" validate:"required,email"`
	}

	// ResetPasswordRequest ResetPasswordRequest
	ResetPasswordRequest struct {
		Password string `json:"password" validate:"required"`
		Token    string `json:"token" validate:"required"`
	}

	SessionContext struct {
		User models.User
		AuthToken models.AuthToken
	}
)

func (g GoogleData) ToUser() models.User {
	return models.User{
		Firstname:         g.GivenName,
		Lastname:          g.FamilyName,
		Email:             g.Email,
		Type:              models.Google,
	}
}

func (r *CreateUserRequest) TransformToUserModel(hp string) models.User {
	return models.User{
		Email:             r.Email,
		Firstname:         r.Firstname,
		Lastname:          r.Lastname,
		HashedPassword:    hp,
	}
}

func (r *CreateDoctorRequest) TransformToUserModel(hp string) models.User {
	return models.User{
		Email:             r.Email,
		Firstname:         r.Firstname,
		Lastname:          r.Lastname,
		HashedPassword:    hp,
		Role:              models.RoleDoctor,
		NoSip:             r.NoSip,
	}
}

func (r *CreatePharmacistRequest) TransformToUserModel(hp string) models.User {
	return models.User{
		Email:             r.Email,
		Firstname:         r.Firstname,
		Lastname:          r.Lastname,
		HashedPassword:    hp,
		Role:              models.RolePharmacist,
		NoSipa:            r.NoSipa,
	}
}

func (r *CreateAdminRequest) TransformToUserModel(hp string) models.User {
	return models.User{
		Email:             r.Email,
		Firstname:         r.Firstname,
		Lastname:          r.Lastname,
		HashedPassword:    hp,
		Role:              models.RolePharmacist,
	}
}