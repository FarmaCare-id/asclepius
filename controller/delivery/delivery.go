package delivery

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
	delivery := app.Group("/delivery")
	delivery.Get("/", c.Shared.Middleware.AuthMiddleware, c.GetAllDeliveryForCurrentUser)
	delivery.Post("/create",  c.Shared.Middleware.AuthMiddleware, c.createDelivery)
}

// All godoc
// @Tags Delivery
// @Summary Get All Delivery
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} common.Response
// @Router /delivery [get]
func (c *Controller) GetAllDeliveryForCurrentUser(ctx *fiber.Ctx) error {
	var (
		res []dto.GetDeliveryResponse
	)	

	context, err := common.CreateContext(ctx)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("getting delivery for user: %s", context.User)

	res, err = c.Service.DeliveryViewService.GetAllDeliveryForCurrentUser(context)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}


// All godoc
// @Tags Delivery
// @Summary Create Delivery
// @Description Put all mandatory parameter
// @Param CreateDeliveryRequest body dto.CreateDeliveryRequest true "CreateDeliveryRequest"
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.CreateDeliveryResponse
// @Failure 200 {array} dto.CreateDeliveryResponse
// @Router /delivery/create [post]
func (c *Controller) createDelivery(ctx *fiber.Ctx) error {
	var req dto.CreateDeliveryRequest

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	context, err := common.CreateContext(ctx)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("creating delivery for user: %s", context.User)

	res, err := c.Service.DeliveryViewService.CreateDelivery(context, req)
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