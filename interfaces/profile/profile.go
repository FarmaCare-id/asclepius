package profile

import (
	"farmacare/application"
	"farmacare/shared"
	"farmacare/shared/dto"
)

type (
	ViewService interface {
		GetUserProfile(user dto.User) dto.User
	}

	viewService struct {
		application application.Holder
		shared      shared.Holder
	}
)

func (v *viewService) GetUserProfile(user dto.User) dto.User {
	user.HashedPassword = ""
	return user
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}