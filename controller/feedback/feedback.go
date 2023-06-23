package feedback

import (
	"farmacare/repository"
	"farmacare/service"
	"farmacare/shared"
	"farmacare/shared/common"
	"farmacare/shared/dto"
	"farmacare/shared/models"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Service service.Holder
	Shared     shared.Holder
	Controller repository.Holder
}

func (c *Controller) Routes(app *fiber.App) {
	feedback := app.Group("/feedback")
	feedback.Get("/", c.Shared.Middleware.AuthMiddleware, c.getAllFeedback)
	feedback.Get("/user/:id", c.Shared.Middleware.AuthMiddleware, c.getAllFeedbackByUserId)
	feedback.Post("/create",  c.Shared.Middleware.AuthMiddleware, c.createFeedback)
	feedback.Get("/category", c.Shared.Middleware.AuthMiddleware, c.getAllFeedbackCategory)
	feedback.Get("/category/:id", c.Shared.Middleware.AuthMiddleware, c.getFeedbackCategoryById)
	feedback.Post("/category/create",  c.Shared.Middleware.AuthMiddleware, c.createFeedbackCategory)
}

// All godoc
// @Tags Feedback
// @Summary Get All Feedback
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} common.Response
// @Router /feedback [get]
func (c *Controller) getAllFeedback(ctx *fiber.Ctx) error {
	var (
		res []dto.GetAllFeedbackResponse
	)

	res, err := c.Service.FeedbackViewService.GetAllFeedback()
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Feedback
// @Summary Get Feedback By User Id
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} common.Response
// @Router /feedback/user/:id [get]
func (c *Controller) getAllFeedbackByUserId(ctx *fiber.Ctx) error {
	var (
		res []dto.GetAllFeedbackByUserIdResponse
	)

	userId := ctx.Params("id")

	c.Shared.Logger.Infof("getting feedback for user: %s", userId)
	
	res, err := c.Service.FeedbackViewService.GetAllFeedbackByUserId(userId)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Feedback
// @Summary Create Feedback
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

	context, err := common.CreateContext(ctx)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("creating feedback for user: %s", context.User)

	res, err := c.Service.FeedbackViewService.CreateFeedback(context, req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Feedback
// @Summary Get All Feedback Category
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 200
// @Router /feedback/category [get]
func (c *Controller) getAllFeedbackCategory(ctx *fiber.Ctx) error {
	var (
		res []dto.GetAllFeedbackCategoryResponse
	)

	res, err := c.Service.FeedbackViewService.GetAllFeedbackCategory()
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Feedback
// @Summary Get Feedback Category By Id
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 200
// @Router /feedback/category/{id} [get]
func (c *Controller) getFeedbackCategoryById(ctx *fiber.Ctx) error {
	var (
		res models.FeedbackCategory
	)

	feedbackId := ctx.Params("id")
	res, err := c.Service.FeedbackViewService.GetFeedbackCategoryById(feedbackId)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Feedback
// @Summary Create Feedback Category
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

	res, err = c.Service.FeedbackViewService.CreateFeedbackCategory(req)
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