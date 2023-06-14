package auth

import (
	"farmacare/shared"
	"farmacare/shared/models"
)

type (
	AuthTokenRepository interface {
		CreateToken(token models.AuthToken) models.AuthToken
		InvalidateToken(userId uint) models.AuthToken
	}

	authTokenRepository struct {
		shared shared.Holder
	}
)

func (s *authTokenRepository) CreateToken(token models.AuthToken) models.AuthToken {
	s.shared.DB.Create(&token)
	return token
}

func (s *authTokenRepository) InvalidateToken(userId uint) models.AuthToken {
	var token models.AuthToken
	s.shared.DB.Where("id = ?", userId).Update("expired_at", "now()")
	return token
}

func NewAuthTokenRepository(holder shared.Holder) (AuthTokenRepository, error) {
	return &authTokenRepository{
		shared: holder,
	}, nil
}