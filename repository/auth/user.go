package auth

import (
	"farmacare/shared"
	"farmacare/shared/models"
)

type (
	UserRepository interface {
		CheckUserExist(email string) (bool, models.User)
		CreateUser(user models.User) error
		CreateDoctor(user models.User) error
		CreatePharmacist(user models.User) error
		EditUser(user models.User) error
		GetUserContext(id uint) models.User
		GetAllUser(preload string) []models.User
	}

	userRepository struct {
		shared shared.Holder
	}
)

func (s *userRepository) CheckUserExist(email string) (bool, models.User) {
	var user models.User
	err := s.shared.DB.First(&user, "email = ?", email).Error
	return err == nil, user
}

func (s *userRepository) CreateUser(user models.User) error {
	user.Role = "user"
	err := s.shared.DB.Create(&user).Error
	return err
}

func (s *userRepository) CreateDoctor(user models.User) error {
	user.Role = "doctor"
	err := s.shared.DB.Create(&user).Error
	return err
}

func (s *userRepository) CreatePharmacist(user models.User) error {
	user.Role = "pharmacist"
	err := s.shared.DB.Create(&user).Error
	return err
}

func (s *userRepository) EditUser(user models.User) error {
	err := s.shared.DB.Save(&user).Error
	return err
}

func (s *userRepository) GetUserContext(id uint) models.User {
	var user models.User
	s.shared.DB.Where("id = ?", id).First(&user)
	return user
}

func (s *userRepository) GetAllUser(preload string) []models.User {
	var users []models.User
	s.shared.DB.Preload(preload).Find(&users)
	return users
}

func NewUserRepository(holder shared.Holder) (UserRepository, error) {
	return &userRepository{
		shared: holder,
	}, nil
}