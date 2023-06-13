package auth

import (
	"farmacare/shared"
	"farmacare/shared/dto"
	"farmacare/shared/models"
)

type (
	Service interface {
		CheckUserExist(email string) (bool, models.User)
		CreateUser(user models.User) error
		CreateDoctor(user models.User) error
		CreatePharmacist(user models.User) error
		CreatePasswordReset(pw dto.PasswordReset) error
		GetResetToken(token string, pw *dto.PasswordReset) error
		RemovePreviousPasswordResetToken(id uint)
		GetUserContext(id uint) models.User
		ListUser(preload string) dto.UserSlice
	}

	service struct {
		shared shared.Holder
	}
)

func (s *service) CheckUserExist(email string) (bool, models.User) {
	var user models.User
	err := s.shared.DB.First(&user, "email = ?", email).Error
	return err == nil, user
}

func (s *service) CreateUser(user models.User) error {
	user.Role = "user"
	err := s.shared.DB.Create(&user).Error
	return err
}

func (s *service) CreateDoctor(user models.User) error {
	user.Role = "doctor"
	err := s.shared.DB.Create(&user).Error
	return err
}

func (s *service) CreatePharmacist(user models.User) error {
	user.Role = "pharmacist"
	err := s.shared.DB.Create(&user).Error
	return err
}

func (s *service) CreatePasswordReset(pw dto.PasswordReset) error {
	err := s.shared.DB.Create(&pw).Error
	return err
}

func (s *service) GetResetToken(token string, pw *dto.PasswordReset) error {
	err := s.shared.DB.Preload("User").First(pw, "token = ?", token).Error
	return err
}

func (s *service) RemovePreviousPasswordResetToken(id uint) {
	var pw dto.PasswordReset
	s.shared.DB.Where("user_id = ?", id).Delete(&pw)
}

func (s *service) GetUserContext(id uint) models.User {
	var user models.User
	s.shared.DB.Where("id = ?", id).First(&user)
	return user
}

func (s *service) ListUser(preload string) dto.UserSlice {
	var users []models.User
	s.shared.DB.Preload(preload).Find(&users)
	return users
}

func NewAuthService(holder shared.Holder) (Service, error) {
	return &service{
		shared: holder,
	}, nil
}