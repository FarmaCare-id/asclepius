package profile

import (
	"farmacare/repository"
	"farmacare/shared"
	"farmacare/shared/dto"
	"farmacare/shared/models"
)

type (
	ViewService interface {
		GetUserProfileRoleUser(ctx dto.SessionContext) (dto.GetUserProfileRoleUserResponse, error)
		GetUserProfileRoleDoctor(ctx dto.SessionContext) (dto.GetUserProfileRoleDoctorResponse, error)
		GetUserProfileRolePharmacist(ctx dto.SessionContext) (dto.GetUserProfileRolePharmacistResponse, error)
		GetUserProfileRoleAdmin(ctx dto.SessionContext) models.User
	}

	viewService struct {
		repository repository.Holder
		shared      shared.Holder
	}
)

func (v *viewService) GetUserProfileRoleUser(ctx dto.SessionContext) (dto.GetUserProfileRoleUserResponse, error) {
	var (
		res dto.GetUserProfileRoleUserResponse
	)

	user := v.repository.UserRepository.GetUserContext(ctx.User.ID)

	res = dto.GetUserProfileRoleUserResponse{
		ID: user.ID,
		Email: user.Email,
		Fullname: user.Fullname,
		Role: user.Role,
		Weight: user.Weight,
		Height: user.Height,
		Age: user.Age,
	}

	return res, nil
}

func (v *viewService) GetUserProfileRoleDoctor(ctx dto.SessionContext) (dto.GetUserProfileRoleDoctorResponse, error) {
	var (
		res dto.GetUserProfileRoleDoctorResponse
	)

	user := v.repository.UserRepository.GetUserContext(ctx.User.ID)

	res = dto.GetUserProfileRoleDoctorResponse{
		ID: user.ID,
		Email: user.Email,
		Fullname: user.Fullname,
		Role: user.Role,
		NoSip: user.NoSip,
		Specialist: user.Specialist,
		Title: user.Title,
	}

	return res, nil
}

func (v *viewService) GetUserProfileRolePharmacist(ctx dto.SessionContext) (dto.GetUserProfileRolePharmacistResponse, error) {
	var (
		res dto.GetUserProfileRolePharmacistResponse
	)

	user := v.repository.UserRepository.GetUserContext(ctx.User.ID)

	res = dto.GetUserProfileRolePharmacistResponse{
		ID: user.ID,
		Email: user.Email,
		Fullname: user.Fullname,
		Role: user.Role,
		NoSipa: user.NoSipa,
		Specialist: user.Specialist,
		Title: user.Title,
	}

	return res, nil
}

func (v *viewService) GetUserProfileRoleAdmin(ctx dto.SessionContext) models.User {
	user := v.repository.UserRepository.GetUserContext(ctx.User.ID)
	
	return user
}

func NewViewService(repository repository.Holder, shared shared.Holder) ViewService {
	return &viewService{
		repository: repository,
		shared:      shared,
	}
}