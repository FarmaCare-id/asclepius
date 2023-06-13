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
	feedback.Post("/create",  c.Shared.Middleware.AuthMiddleware, c.createFeedback)
	feedback.Get("/category", c.Shared.Middleware.AuthMiddleware, c.getFeedbackCategory)
	feedback.Post("/category/create",  c.Shared.Middleware.AuthMiddleware, c.createFeedbackCategory)
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
// @Summary Get feedback category
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.FeedbackCategorySlice
// @Failure 200 {object} dto.FeedbackCategorySlice
// @Router /feedback/category [get]
func (c *Controller) getFeedbackCategory(ctx *fiber.Ctx) error {
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
	var (
		req dto.CreateFeedbackRequest
		res dto.CreateFeedbackResponse
		usr dto.User
	)

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("create feedback with payload: %s", req)

	res, err = c.Interfaces.FeedbackViewService.CreateFeedback(req, usr)
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