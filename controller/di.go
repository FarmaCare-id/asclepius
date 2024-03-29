package controller

import (
	"farmacare/controller/appointment"
	"farmacare/controller/auth"
	"farmacare/controller/community"
	"farmacare/controller/delivery"
	"farmacare/controller/feedback"
	"farmacare/controller/healthcheck"
	"farmacare/controller/management"
	"farmacare/controller/profile"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	Healthcheck  healthcheck.Controller
	Auth         auth.Controller
	Profile 	 profile.Controller
	Feedback	 feedback.Controller
	Management   management.Controller
	Delivery     delivery.Controller
	Community	 community.Controller
	Appointment  appointment.Controller
}

func Register(container *dig.Container) error {
	if err := container.Provide(healthcheck.NewController); err != nil {
		return errors.Wrap(err, "failed to provide healthcheck controller")
	}

	if err := container.Provide(auth.NewController); err != nil {
		return errors.Wrap(err, "failed to provide auth controller")
	}

	if err := container.Provide(profile.NewController); err != nil {
		return errors.Wrap(err, "failed to provide profile controller")
	}

	if err := container.Provide(feedback.NewController); err != nil {
		return errors.Wrap(err, "failed to provide feedback controller")
	}
	
	if err := container.Provide(management.NewController); err != nil {
		return errors.Wrap(err, "failed to provide management controller")
	}

	if err := container.Provide(delivery.NewController); err != nil {
		return errors.Wrap(err, "failed to provide delivery controller")
	}

	if err := container.Provide(community.NewController); err != nil {
		return errors.Wrap(err, "failed to provide community controller")
	}

	if err := container.Provide(appointment.NewController); err != nil {
		return errors.Wrap(err, "failed to provide appointment controller")
	}

	return nil
}

func Routes(app *fiber.App, controller Holder) {
	controller.Healthcheck.Routes(app)
	controller.Auth.Routes(app)
	controller.Profile.Routes(app)
	controller.Feedback.Routes(app)
	controller.Management.Routes(app)
	controller.Delivery.Routes(app)
	controller.Community.Routes(app)
	controller.Appointment.Routes(app)
}