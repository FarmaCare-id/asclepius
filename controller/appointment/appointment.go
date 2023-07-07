package appointment

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
	appointment := app.Group("/appointment")
	appointment.Get("/user",  c.Shared.Middleware.AuthMiddleware, c.getAllAppointmentForCurrentUser)
	appointment.Get("/user/:id", c.Shared.Middleware.AuthMiddleware, c.getAllAppointmentByUserID)
	appointment.Post("/create",  c.Shared.Middleware.AuthMiddleware, c.createAppointment)
}

// All godoc
// @Tags Appointment
// @Summary Get Appointment By User Id
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} common.Response
// @Router /appointment/user/:id [get]
func (c *Controller) getAllAppointmentByUserID(ctx *fiber.Ctx) error {
	var (
		res []dto.GetAppointmentResponse
	)

	userId := ctx.Params("id")

	c.Shared.Logger.Infof("getting appointment for user: %s", userId)
	
	res, err := c.Service.AppointmentViewService.GetAllAppointmentByUserID(userId)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Appointment
// @Summary Create Appointment
// @Description Put all mandatory parameter
// @Param CreateAppointmentRequest body dto.CreateAppointmentRequest true "CreateAppointmentRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.CreateAppointmentResponse
// @Failure 200 {object} dto.CreateAppointmentResponse
// @Router /appointment/create [post]
func (c *Controller) createAppointment(ctx *fiber.Ctx) error {
	var req dto.CreateAppointmentRequest

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	context, err := common.CreateContext(ctx)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("creating appointment for user: %s", context.User)

	res, err := c.Service.AppointmentViewService.CreateAppointment(context, req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Appointment
// @Summary Get All Appointment For Current User
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400 {object} common.Response
// @Router /appointment/user [get]
func (c *Controller) getAllAppointmentForCurrentUser(ctx *fiber.Ctx) error {
	var (
		res []dto.GetAppointmentResponse
	)

	context, err := common.CreateContext(ctx)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("getting appointment for user: %s", context.User)

	res, err = c.Service.AppointmentViewService.GetAllAppointmentForCurrentUser(context)
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