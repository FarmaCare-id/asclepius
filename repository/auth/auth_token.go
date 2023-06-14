package auth

import (
	"farmacare/shared"
	"farmacare/shared/models"
)

type (
	AuthTokenRepository interface {
		CreateToken(token models.AuthToken) models.AuthToken
		InvalidateToken(token models.AuthToken) models.AuthToken
	}

	authTokenRepository struct {
		shared shared.Holder
	}
)

func (s *authTokenRepository) CreateToken(token models.AuthToken) models.AuthToken {
	s.shared.DB.Create(&token)
	return token
}

func (s *authTokenRepository) InvalidateToken(token models.AuthToken) models.AuthToken {
	s.shared.DB.Where("token = ?", &token.Token).Delete(&token)
	return token
}

func NewAuthTokenRepository(holder shared.Holder) (AuthTokenRepository, error) {
	return &authTokenRepository{
		shared: holder,
	}, nil
}