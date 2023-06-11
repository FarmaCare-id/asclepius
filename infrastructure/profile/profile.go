package profile

import (
	"farmacare/application"
	"farmacare/interfaces"
	"farmacare/shared"
	"farmacare/shared/common"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Interfaces  interfaces.Holder
	Shared      shared.Holder
	Application application.Holder
}

func (c *Controller) Routes(app *fiber.App) {
	profile := app.Group("/profile")
	profile.Get("/", c.Shared.Middleware.AuthMiddleware, c.userProfile)
}

// All godoc
// @Tags Profile
// @Summary Get user profile
// @Param Authorization header string true "Authorization"
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200
// @Router /profile [get]
func (c *Controller) userProfile(ctx *fiber.Ctx) error {
	user := common.CreateUserContext(ctx, c.Application.AuthService)

	user = c.Interfaces.ProfileViewService.GetUserProfile(user)

	return common.DoCommonSuccessResponse(ctx, user)
}

func NewController(interfaces interfaces.Holder, shared shared.Holder, application application.Holder) Controller {
	return Controller{
		Interfaces:  interfaces,
		Shared:      shared,
		Application: application,
	}
}