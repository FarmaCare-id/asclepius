package service

import (
	"farmacare/service/auth"
	"farmacare/service/delivery"
	"farmacare/service/feedback"
	"farmacare/service/healthcheck"
	"farmacare/service/management"
	"farmacare/service/profile"

	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	HealthcheckViewService  healthcheck.ViewService
	AuthViewService         auth.ViewService
	FeedbackViewService		feedback.ViewService
	ProfileViewService		profile.ViewService
	ManagementViewService	management.ViewService
	DeliveryViewService		delivery.ViewService
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

	if err := container.Provide(profile.NewViewService); err != nil {
		return errors.Wrap(err, "Failed to provide profile view repository")
	}

	if err := container.Provide(management.NewViewService); err != nil {
		return errors.Wrap(err, "Failed to provide management view repository")
	}

	if err := container.Provide(delivery.NewViewService); err != nil {
		return errors.Wrap(err, "Failed to provide delivery view repository")
	}

	return nil
}