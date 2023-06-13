package dto

import (
	"farmacare/shared/models"
	"time"
)

type (
	UserSlice []models.User

	PasswordReset struct {
		ID        uint `gorm:"primaryKey;autoIncrement"`
		UserID    uint
		User      models.User   `gorm:"onDelete:CASCADE"`
		Token     string `gorm:"column:token"`
		CreatedAt time.Time
		Valid     time.Time
	}	

	// CreateUserRequest CreateUserRequest
	CreateUserRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Fullname string `json:"fullname" validate:"required"`
	}

	// CreateUserResponse CreateUserResponse
	CreateUserResponse struct {
		Email    string `json:"email"`
		Fullname string `json:"fullname"`
		Role     string `json:"role"`
	}

	// CreateDoctorRequest CreateDoctorRequest
	CreateDoctorRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Fullname string `json:"fullname" validate:"required"`
		NoSip    string `json:"no_sip" validate:"required"`
	}

	// CreateDoctorResponse CreateDoctorResponse
	CreateDoctorResponse struct {
		Email    string `json:"email"`
		Fullname string `json:"fullname"`
		Role     string `json:"role"`
		NoSip    string `json:"no_sip"`
	}

	// CreatePharmacistRequest CreatePharmacistRequest
	CreatePharmacistRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Fullname string `json:"fullname" validate:"required"`
		NoSipa   string `json:"no_sipa" validate:"required"`
	}

	// CreatePharmacistResponse CreatePharmacistResponse
	CreatePharmacistResponse struct {
		Email    string `json:"email"`
		Fullname string `json:"fullname"`
		Role     string `json:"role"`
		NoSipa   string `json:"no_sipa"`
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
	}
)

func (g GoogleData) ToUser() models.User {
	return models.User{
		Fullname:          g.Name,
		Email:             g.Email,
		Type:              models.Google,
	}
}

func (r *CreateUserRequest) TransformToUserModel(hp string) models.User {
	return models.User{
		Email:             r.Email,
		Fullname:          r.Fullname,
		HashedPassword:    hp,
	}
}

func (r *CreateDoctorRequest) TransformToUserModel(hp string) models.User {
	return models.User{
		Email:             r.Email,
		Fullname:          r.Fullname,
		HashedPassword:    hp,
		Role:              models.RoleDoctor,
		NoSip:             r.NoSip,
	}
}

func (r *CreatePharmacistRequest) TransformToUserModel(hp string) models.User {
	return models.User{
		Email:             r.Email,
		Fullname:          r.Fullname,
		HashedPassword:    hp,
		Role:              models.RolePharmacist,
		NoSipa:            r.NoSipa,
	}
}