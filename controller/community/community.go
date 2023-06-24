package community

import (
	"farmacare/repository"
	"farmacare/service"
	"farmacare/shared"
	"farmacare/shared/common"
	"farmacare/shared/dto"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Service service.Holder
	Shared     shared.Holder
	Controller repository.Holder
}

func (c *Controller) Routes(app *fiber.App) {
	community := app.Group("/community")
	community.Post("/forum/create",  c.Shared.Middleware.AuthMiddleware, c.createForum)
}

// All godoc
// @Tags Community
// @Summary Create Forum
// @Description Put all mandatory parameter
// @Param CreateForumRequest body CreateForumRequest true "CreateForumRequest"
// @Accept  json
// @Produce  json
// @Success 200 {array} CreateForumResponse
// @Failure 200 {array} CreateForumResponse
// @Router /community/create [post]
func (c *Controller) createForum(ctx *fiber.Ctx) error {
	var (
		req dto.CreateForumRequest
	)

	context, err := common.CreateContext(ctx)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("creating forum for user: %s", context.User)

	if err := ctx.BodyParser(&req); err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	res, err := c.Service.CommunityViewService.CreateForum(context, req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

func NewController(service service.Holder, shared shared.Holder, repository repository.Holder) Controller {
	return Controller{
		Service:  service,
		Shared:      shared,
		Controller: repository,
	}
}