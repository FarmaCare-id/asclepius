package profile

import (
	"farmacare/repository"
	"farmacare/shared"
	"farmacare/shared/dto"
)

type (
	ViewService interface {
		GetUserProfileRoleUser(ctx dto.SessionContext) (dto.GetUserProfileRoleUserResponse, error)
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
		Role: user.Role,
		Fullname: user.Fullname,
		Weight: user.Weight,
		Height: user.Height,
		Age: user.Age,
	}

	return res, nil
}

func NewViewService(repository repository.Holder, shared shared.Holder) ViewService {
	return &viewService{
		repository: repository,
		shared:      shared,
	}
}