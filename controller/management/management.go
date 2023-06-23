package management

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
	management := app.Group("/management")
	management.Get("/drug",  c.Shared.Middleware.AuthMiddleware, c.getAllDrug)
	management.Post("/drug/create",  c.Shared.Middleware.AuthMiddleware, c.createDrug)
}

// All godoc
// @Tags Management
// @Summary Get All Drug
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} common.Response
// @Router /management/drug [get]
func (c *Controller) getAllDrug(ctx *fiber.Ctx) error {
	var (
		res []dto.GetAllDrugResponse
	)

	res, err := c.Service.ManagementViewService.GetAllDrug()
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Management
// @Summary Create Drug
// @Description Put all mandatory parameter
// @Param CreateDrugRequest body dto.CreateDrugRequest true "CreateDrugRequest"
// @Accept  json
// @Produce  json
// @Success 200 {array} dto.CreateDrugResponse
// @Failure 200 {array} dto.CreateDrugResponse
// @Router /management/drug/create [post]
func (c *Controller) createDrug(ctx *fiber.Ctx) error {
	var req dto.CreateDrugRequest

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	context, err := common.CreateContext(ctx)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("creating drug for user: %s", context.User)

	res, err := c.Service.ManagementViewService.CreateDrug(context, req)
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