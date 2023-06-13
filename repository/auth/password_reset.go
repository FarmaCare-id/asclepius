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

	passwordResetrepository struct {
		shared shared.Holder
	}
)

func (s *passwordResetrepository) CreatePasswordReset(pw models.PasswordReset) error {
	err := s.shared.DB.Create(&pw).Error
	return err
}

func (s *passwordResetrepository) GetResetToken(token string, pw *models.PasswordReset) error {
	err := s.shared.DB.Preload("User").First(pw, "token = ?", token).Error
	return err
}

func (s *passwordResetrepository) RemovePreviousPasswordResetToken(id uint) {
	var pw models.PasswordReset
	s.shared.DB.Where("user_id = ?", id).Delete(&pw)
}

func NewPasswordResetRepository(holder shared.Holder) (PasswordResetRepository, error) {
	return &passwordResetrepository{
		shared: holder,
	}, nil
}