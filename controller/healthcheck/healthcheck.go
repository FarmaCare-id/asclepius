package healthcheck

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

// All godoc
// @Tags Healthcheck
// @Summary Check System Status
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.Status
// @Failure 200 {array} dto.Status
// @Router /healthcheck [get]
func (c *Controller) healthcheck(ctx *fiber.Ctx) error {
	c.Shared.Logger.Println("checking server status")
	data, _ := c.Interfaces.HealthcheckViewService.SystemHealthcheck()
	return ctx.Status(fiber.StatusOK).JSON(data)
}

func NewController(service service.Holder, shared shared.Holder, repository repository.Holder) Controller {
	return Controller{
		Interfaces:  service,
		Shared:      shared,
		Application: repository,
	}
}