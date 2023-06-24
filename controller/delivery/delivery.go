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
	delivery.Post("/create",  c.Shared.Middleware.AuthMiddleware, c.createDelivery)
}

// All godoc
// @Tags Delivery
// @Summary Create Delivery
// @Description Put all mandatory parameter
// @Param CreateDeliveryRequest body CreateDeliveryRequest true "CreateDeliveryRequest"
// @Accept  json
// @Produce  json
// @Success 200 {array} CreateDeliveryResponse
// @Failure 200 {array} CreateDeliveryResponse
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