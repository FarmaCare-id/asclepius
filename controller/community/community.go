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
	community.Post("/forum/:id/comment/create",  c.Shared.Middleware.AuthMiddleware, c.createForumComment)
	community.Get("/forum/:id/comment/", c.Shared.Middleware.AuthMiddleware, c.getAllForumCommentByForumID)
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

	forumId := ctx.Params("id")

	context, err := common.CreateContext(ctx)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("creating forum comment for user: %s", context.User)
	c.Shared.Logger.Infof("creating forum comment for forum: %s", forumId)

	if err := ctx.BodyParser(&req); err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	res, err := c.Service.CommunityViewService.CreateForumComment(context, req, forumId)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Community
// @Summary Get All Forum Comment By Forum ID
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} common.Response
// @Router /community/forum/:id/comment/ [get]
func (c *Controller) getAllForumCommentByForumID(ctx *fiber.Ctx) error {
	var (
		res []dto.GetForumCommentResponse
	)

	forumId := ctx.Params("id")

	c.Shared.Logger.Infof("getting forum comment for forum: %s", forumId)

	res, err := c.Service.CommunityViewService.GetAllForumCommentByForumID(forumId)
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