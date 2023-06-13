package service

import (
	"farmacare/service/auth"
	"farmacare/service/feedback"
	"farmacare/service/healthcheck"
	"farmacare/service/profile"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	HealthcheckViewService  healthcheck.ViewService
	AuthViewService         auth.ViewService
	ProfileViewService		profile.ViewService
	FeedbackViewService		feedback.ViewService
}

func Register(container *dig.Container) error {
	if err := container.Provide(healthcheck.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide healthcheck view service")
	}

	if err := container.Provide(auth.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide auth view service")
	}

	if err := container.Provide(profile.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide profile view service")
	}

	if err := container.Provide(feedback.NewViewService); err != nil {
		return errors.Wrap(err, "failed to provide profile view service")
	}

	return nil
}