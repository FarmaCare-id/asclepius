package application

import (
	"farmacare/application/auth"
	"farmacare/application/feedback"
	"farmacare/application/healthcheck"
	"farmacare/application/profile"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	HealthcheckService  healthcheck.Service
	AuthService         auth.Service
	ProfileService		profile.Service
	FeedbackService     feedback.Service
}

func Register(container *dig.Container) error {
	if err := container.Provide(healthcheck.NewHealthcheckService); err != nil {
		return errors.Wrap(err, "failed to provide healthcheck service")
	}

	if err := container.Provide(auth.NewAuthService); err != nil {
		return errors.Wrap(err, "failed to provide auth service")
	}

	if err := container.Provide(profile.NewProfileService); err != nil {
		return errors.Wrap(err, "failed to provide profile service")
	}

	if err := container.Provide(feedback.NewFeedbackService); err != nil {
		return errors.Wrap(err, "failed to provide feedback service")
	}
	
	return nil
}