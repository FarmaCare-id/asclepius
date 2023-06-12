package profile

import (
	"errors"
	"farmacare/application"
	"farmacare/shared"
	"farmacare/shared/dto"

	"golang.org/x/crypto/bcrypt"
)

type (
	ViewService interface {
		GetUserProfile(user dto.User) dto.User
		EditUserProfile(req dto.EditUserRequest, user dto.User) (dto.EditUserResponse, error)
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

func (v *viewService) EditUserProfile(req dto.EditUserRequest, user dto.User) (dto.EditUserResponse, error) {
	var (
		res dto.EditUserResponse
	)

	isUserExist, user := v.application.AuthService.CheckUserExist(user.Email)
	if !isUserExist {
		return res, errors.New("no user found for given email")
	}

	if req.Fullname != "" {
		user.Fullname = req.Fullname
	}

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
		if err != nil {
			return res, err
		}
		user.HashedPassword = string(hashedPassword)
	}

	err := v.application.ProfileService.EditUserProfile(user)
	if err != nil {
		return res, err
	}

	res = dto.EditUserResponse{
		Email:    user.Email,
		Fullname: req.Fullname,
	}

	return res, nil
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}