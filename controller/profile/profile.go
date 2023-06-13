package profile

import (
	"farmacare/repository"
	"farmacare/service"
	"farmacare/shared"
	"farmacare/shared/common"
	"farmacare/shared/dto"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Service  service.Holder
	Shared      shared.Holder
	Application repository.Holder
}

func (c *Controller) Routes(app *fiber.App) {
	profile := app.Group("/profile")
	profile.Get("/", c.Shared.Middleware.AuthMiddleware, c.userProfile)
	profile.Put("/edit", c.Shared.Middleware.AuthMiddleware, c.editProfile)
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
	context := common.CreateContext(ctx)

	user := c.Service.AuthViewService.GetUserCredential(context)

	return common.DoCommonSuccessResponse(ctx, user)
}

// All godoc
// @Tags Profile
// @Summary Edit user profile
// @Description Put all mandatory parameter
// @Param Authorization header string true "Authorization"
// @Param EditUserProfileRequest body dto.EditUserProfileRequest true "EditUserProfileRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.EditUserProfileResponse
// @Failure 200 {object} dto.EditUserProfileResponse
// @Router /profile/edit [put]
func (c *Controller) editProfile(ctx *fiber.Ctx) error {
	var (
		req dto.EditUserProfileRequest
		res dto.EditUserProfileResponse
	)

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("edit user with payload: %s", req)

	context := common.CreateContext(ctx)

	res, err = c.Service.ProfileViewService.EditUserProfile(req, context)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

func NewController(service service.Holder, shared shared.Holder, repository repository.Holder) Controller {
	return Controller{
		Service:  service,
		Shared:      shared,
		Application: repository,
	}
}