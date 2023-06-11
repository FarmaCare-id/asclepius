package infrastructure

import (
	"farmacare/infrastructure/auth"
	"farmacare/infrastructure/healthcheck"
	"farmacare/infrastructure/profile"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type Holder struct {
	dig.In
	Healthcheck  healthcheck.Controller
	Auth         auth.Controller
	Profile 	 profile.Controller
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

	return nil
}

func Routes(app *fiber.App, controller Holder) {
	controller.Healthcheck.Routes(app)
	controller.Auth.Routes(app)
	controller.Profile.Routes(app)
}