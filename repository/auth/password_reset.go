package auth

import (
	"farmacare/shared"
	"farmacare/shared/models"
)

type (
	PasswordResetRepository interface {
		CreatePasswordReset(pw models.PasswordReset) error
		GetResetToken(token string, pw *models.PasswordReset) error
		RemovePreviousPasswordResetToken(id uint)
	}

	passwordResetRepository struct {
		shared shared.Holder
	}
)

func (s *passwordResetRepository) CreatePasswordReset(pw models.PasswordReset) error {
	err := s.shared.DB.Create(&pw).Error
	return err
}

func (s *passwordResetRepository) GetResetToken(token string, pw *models.PasswordReset) error {
	err := s.shared.DB.Preload("User").First(pw, "token = ?", token).Error
	return err
}

func (s *passwordResetRepository) RemovePreviousPasswordResetToken(id uint) {
	var pw models.PasswordReset
	s.shared.DB.Where("user_id = ?", id).Delete(&pw)
}

func NewPasswordResetRepository(holder shared.Holder) (PasswordResetRepository, error) {
	return &passwordResetRepository{
		shared: holder,
	}, nil
}