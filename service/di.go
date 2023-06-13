package service

import (
	"farmacare/service/auth"
	"farmacare/service/feedback"
	"farmacare/service/healthcheck"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	HealthcheckViewService  healthcheck.ViewService
	AuthViewService         auth.ViewService
	FeedbackViewService		feedback.ViewService
}

func Register(container *dig.Container) error {
	if err := container.Provide(healthcheck.NewViewService); err != nil {
		return errors.Wrap(err, "Failed to provide healthcheck view service")
	}

	if err := container.Provide(auth.NewViewService); err != nil {
		return errors.Wrap(err, "Failed to provide auth view service")
	}

	if err := container.Provide(feedback.NewViewService); err != nil {
		return errors.Wrap(err, "Failed to provide profile view repository")
	}

	return nil
}