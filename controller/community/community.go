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
	community.Get("/forum", c.Shared.Middleware.AuthMiddleware, c.getAllForum)
	community.Post("/forum/create",  c.Shared.Middleware.AuthMiddleware, c.createForum)
	community.Post("/forum/comment/create",  c.Shared.Middleware.AuthMiddleware, c.createForumComment)
}

// All godoc
// @Tags Community
// @Summary Get All Forum
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} common.Response
// @Router /community/forum [get]
func (c *Controller) getAllForum(ctx *fiber.Ctx) error {
	var (
		res []dto.GetForumResponse
	)

	res, err := c.Service.CommunityViewService.GetAllForum()
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Community
// @Summary Create Forum
// @Description Put all mandatory parameter
// @Param CreateForumRequest body dto.CreateForumRequest true "CreateForumRequest"
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.CreateForumResponse
// @Failure 200 {array} dto.CreateForumResponse
// @Router /community/forum/create [post]
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

// All godoc
// @Tags Community
// @Summary Create Forum Comment
// @Description Put all mandatory parameter
// @Param CreateForumCommentRequest body dto.CreateForumCommentRequest true "CreateForumCommentRequest"
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.CreateForumCommentResponse
// @Failure 200 {array} dto.CreateForumCommentResponse
// @Router /community/forum/comment/create [post]
func (c *Controller) createForumComment(ctx *fiber.Ctx) error {
	var (
		req dto.CreateForumCommentRequest
	)

	context, err := common.CreateContext(ctx)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("creating forum comment for user: %s", context.User)

	if err := ctx.BodyParser(&req); err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	res, err := c.Service.CommunityViewService.CreateForumComment(context, req)
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