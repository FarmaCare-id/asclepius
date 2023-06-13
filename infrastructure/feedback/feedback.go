package feedback

import (
	"farmacare/application"
	"farmacare/interfaces"
	"farmacare/shared"
	"farmacare/shared/common"
	"farmacare/shared/dto"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Interfaces interfaces.Holder
	Shared     shared.Holder
	Application application.Holder
}

func (c *Controller) Routes(app *fiber.App) {
	feedback := app.Group("/feedback")
	feedback.Get("/", c.Shared.Middleware.AuthMiddleware, c.listFeedback)
	feedback.Post("/create",  c.Shared.Middleware.AuthMiddleware, c.createFeedback)
	feedback.Get("/category", c.Shared.Middleware.AuthMiddleware, c.listFeedbackCategory)
	feedback.Get("/category/:id", c.Shared.Middleware.AuthMiddleware, c.getFeedbackCategoryById)
	feedback.Post("/category/create",  c.Shared.Middleware.AuthMiddleware, c.createFeedbackCategory)
}

// All godoc
// @Tags Feedback
// @Summary List feedback
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.FeedbackSlice
// @Failure 200 {object} dto.FeedbackSlice
// @Router /feedback [get]
func (c *Controller) listFeedback(ctx *fiber.Ctx) error {
	var (
		res []dto.GetAllFeedbackResponse
	)

	res, err := c.Interfaces.FeedbackViewService.ListFeedback()
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

func (c *Controller) getFeedbackCategoryById(ctx *fiber.Ctx) error {
	var (
		res dto.FeedbackCategory
	)

	feedbackId := ctx.Params("id")
	res, err := c.Interfaces.FeedbackViewService.GetFeedbackCategoryById(feedbackId)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Feedback
// @Summary Create feedback category
// @Description Put all mandatory parameter
// @Param CreateFeedbackCategoryRequest body dto.CreateFeedbackCategoryRequest true "CreateFeedbackCategoryRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.CreateFeedbackCategoryResponse
// @Failure 200 {object} dto.CreateFeedbackCategoryResponse
// @Router /feedback/category/create [post]
func (c *Controller) createFeedbackCategory(ctx *fiber.Ctx) error {
	var (
		req dto.CreateFeedbackCategoryRequest
		res dto.CreateFeedbackCategoryResponse
	)

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("create feedback category with payload: %s", req)

	res, err = c.Interfaces.FeedbackViewService.CreateFeedbackCategory(req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Feedback
// @Summary List feedback category
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.FeedbackCategorySlice
// @Failure 200 {object} dto.FeedbackCategorySlice
// @Router /feedback/category [get]
func (c *Controller) listFeedbackCategory(ctx *fiber.Ctx) error {
	var (
		res dto.FeedbackCategorySlice
	)

	res, err := c.Interfaces.FeedbackViewService.ListFeedbackCategory()
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Feedback
// @Summary Create feedback
// @Description Put all mandatory parameter
// @Param CreateFeedbackRequest body dto.CreateFeedbackRequest true "CreateFeedbackRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.CreateFeedbackResponse
// @Failure 200 {object} dto.CreateFeedbackResponse
// @Router /feedback/create [post]
func (c *Controller) createFeedback(ctx *fiber.Ctx) error {
	var req dto.CreateFeedbackRequest

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	context := common.CreateContext(ctx)

	c.Shared.Logger.Infof("creating feedback for user: %s", context.User)

	res, err := c.Interfaces.FeedbackViewService.CreateFeedback(context, req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

func NewController(interfaces interfaces.Holder, shared shared.Holder, application application.Holder) Controller {
	return Controller{
		Interfaces:  interfaces,
		Shared:      shared,
		Application: application,
	}
}