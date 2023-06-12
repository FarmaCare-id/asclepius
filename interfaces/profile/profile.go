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
		EditUserProfile(req dto.EditUserProfileRequest, user dto.User) (dto.EditUserProfileResponse, error)
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

func (v *viewService) EditUserProfile(req dto.EditUserProfileRequest, user dto.User) (dto.EditUserProfileResponse, error) {
	var (
		res dto.EditUserProfileResponse
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

	if req.Weight != 0 {
		user.Weight = req.Weight
	}

	if req.Height != 0 {
		user.Height = req.Height
	}

	if req.Age != 0 {
		user.Age = req.Age
	}

	if req.NoSip != "" {
		user.NoSip = req.NoSip
	}

	if req.NoSipa != "" {
		user.NoSipa = req.NoSipa
	}

	if req.Specialist != "" {
		user.Specialist = req.Specialist
	}

	if req.Title != "" {
		user.Title = req.Title
	}

	err := v.application.ProfileService.EditUserProfile(user)
	if err != nil {
		return res, err
	}

	res = dto.EditUserProfileResponse{
		Fullname: req.Fullname,
		Weight:   req.Weight,
		Height:   req.Height,
		Age:      req.Age,
		NoSip:    req.NoSip,
		NoSipa:   req.NoSipa,
		Specialist: req.Specialist,
		Title: req.Title,
	}

	return res, nil
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}