package repository

import (
	"farmacare/repository/auth"
	"farmacare/repository/feedback"
	"farmacare/repository/healthcheck"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	HealthcheckRepository   healthcheck.Repository
	UserRepository          auth.UserRepository
	AuthTokenRepository 	auth.AuthTokenRepository
	PasswordResetRepository auth.PasswordResetRepository
	FeedbackRepository      feedback.Repository
}

func Register(container *dig.Container) error {
	if err := container.Provide(healthcheck.HealthcheckRepository); err != nil {
		return errors.Wrap(err, "Failed to provide healthcheck repository")
	}

	if err := container.Provide(auth.NewUserRepository); err != nil {
		return errors.Wrap(err, "Failed to provide auth repository")
	}

	if err := container.Provide(auth.NewAuthTokenRepository); err != nil {
		return errors.Wrap(err, "Failed to provide token repository")
	}

	if err := container.Provide(auth.NewPasswordResetRepository); err != nil {
		return errors.Wrap(err, "Failed to provide password_reset repository")
	}

	if err := container.Provide(feedback.FeedbackRepository); err != nil {
		return errors.Wrap(err, "Failed to provide feedback repository")
	}
	
	return nil
}