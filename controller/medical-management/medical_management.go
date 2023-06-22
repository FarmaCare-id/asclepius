package medical_management

import (
	"farmacare/repository"
	"farmacare/service"
	"farmacare/shared"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Interfaces service.Holder
	Shared     shared.Holder
	Application repository.Holder
}

func (c *Controller) Routes(app *fiber.App) {
	healthcheck := app.Group("/healthcheck")
	healthcheck.Get("/",  c.Shared.Middleware.AdminMiddleware, c.healthcheck)
}