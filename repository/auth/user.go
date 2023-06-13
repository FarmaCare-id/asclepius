package auth

import (
	"farmacare/shared"
	"farmacare/shared/dto"
	"farmacare/shared/models"
)

type (
	Repository interface {
		CheckUserExist(email string) (bool, models.User)
		CreateUser(user models.User) error
		CreateDoctor(user models.User) error
		CreatePharmacist(user models.User) error
		EditUser(user models.User) error
		CreatePasswordReset(pw models.PasswordReset) error
		GetResetToken(token string, pw *models.PasswordReset) error
		RemovePreviousPasswordResetToken(id uint)
		GetUserContext(id uint) models.User
		ListUser(preload string) dto.UserSlice
	}

	repository struct {
		shared shared.Holder
	}
)

func (s *repository) CheckUserExist(email string) (bool, models.User) {
	var user models.User
	err := s.shared.DB.First(&user, "email = ?", email).Error
	return err == nil, user
}

func (s *repository) CreateUser(user models.User) error {
	user.Role = "user"
	err := s.shared.DB.Create(&user).Error
	return err
}

func (s *repository) CreateDoctor(user models.User) error {
	user.Role = "doctor"
	err := s.shared.DB.Create(&user).Error
	return err
}

func (s *repository) CreatePharmacist(user models.User) error {
	user.Role = "pharmacist"
	err := s.shared.DB.Create(&user).Error
	return err
}

func (s *repository) EditUser(user models.User) error {
	err := s.shared.DB.Save(&user).Error
	return err
}

func (s *repository) CreatePasswordReset(pw models.PasswordReset) error {
	err := s.shared.DB.Create(&pw).Error
	return err
}

func (s *repository) GetResetToken(token string, pw *models.PasswordReset) error {
	err := s.shared.DB.Preload("User").First(pw, "token = ?", token).Error
	return err
}

func (s *repository) RemovePreviousPasswordResetToken(id uint) {
	var pw models.PasswordReset
	s.shared.DB.Where("user_id = ?", id).Delete(&pw)
}

func (s *repository) GetUserContext(id uint) models.User {
	var user models.User
	s.shared.DB.Where("id = ?", id).First(&user)
	return user
}

func (s *repository) ListUser(preload string) dto.UserSlice {
	var users []models.User
	s.shared.DB.Preload(preload).Find(&users)
	return users
}

func AuthRepository(holder shared.Holder) (Repository, error) {
	return &repository{
		shared: holder,
	}, nil
}