package auth

import (
	"farmacare/application"
	"farmacare/interfaces"
	"farmacare/shared"
	"farmacare/shared/common"
	"farmacare/shared/dto"

	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	Interfaces  interfaces.Holder
	Shared      shared.Holder
	Application application.Holder
}

func (c *Controller) Routes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", c.register)
	auth.Post("/register/doctor", c.registerDoctor)
	auth.Post("/register/pharmacist", c.registerPharmacist)
	auth.Post("/login", c.login)
	auth.Post("/login-google", c.loginGoogle)
	auth.Put("/edit", c.Shared.Middleware.AuthMiddleware, c.edit)
	auth.Post("/forgot-password", c.forgotPassword)
	auth.Post("/reset-password", c.resetPassword)
	auth.Get("/credential", c.Shared.Middleware.AuthMiddleware, c.userCredential)
}

// All godoc
// @Tags Auth
// @Summary Register new user
// @Description Put all mandatory parameter
// @Param CreateUserRequest body dto.CreateUserRequest true "CreateUserRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.CreateUserResponse
// @Failure 200 {object} dto.CreateUserResponse
// @Router /auth/register [post]
func (c *Controller) register(ctx *fiber.Ctx) error {
	var (
		req dto.CreateUserRequest
		res dto.CreateUserResponse
	)

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("register user with payload: %s", req)

	res, err = c.Interfaces.AuthViewService.RegisterUser(req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Auth
// @Summary Register new doctor
// @Description Put all mandatory parameter
// @Param CreateDoctorRequest body dto.CreateDoctorRequest true "CreateDoctorRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.CreateDoctorResponse
// @Failure 200 {object} dto.CreateDoctorResponse
// @Router /auth/register/doctor [post]
func (c *Controller) registerDoctor(ctx *fiber.Ctx) error {
	var (
		req dto.CreateDoctorRequest
		res dto.CreateDoctorResponse
	)

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("register doctor with payload: %s", req)

	res, err = c.Interfaces.AuthViewService.RegisterDoctor(req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Auth
// @Summary Register new pharmacist
// @Description Put all mandatory parameter
// @Param CreatePharmacistRequest body dto.CreatePharmacistRequest true "CreatePharmacistRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.CreatePharmacistResponse
// @Failure 200 {object} dto.CreatePharmacistResponse
// @Router /auth/register/pharmacist [post]
func (c *Controller) registerPharmacist(ctx *fiber.Ctx) error {
	var (
		req dto.CreatePharmacistRequest
		res dto.CreatePharmacistResponse
	)

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("register pharmacist with payload: %s", req)

	res, err = c.Interfaces.AuthViewService.RegisterPharmacist(req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Auth
// @Summary Login user
// @Description Put all mandatory parameter
// @Param GoogleLoginRequest body dto.GoogleLoginRequest true "GoogleLoginRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.LoginResponse
// @Failure 200 {object} dto.LoginResponse
// @Router /auth/login-google [post]
func (c *Controller) loginGoogle(ctx *fiber.Ctx) error {
	var (
		req dto.GoogleLoginRequest
		res dto.LoginResponse
	)

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("login user google with payload: %s", req)

	res, err = c.Interfaces.AuthViewService.GoogleLogin(req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Auth
// @Summary Login user google
// @Description Put all mandatory parameter
// @Param LoginRequest body dto.LoginRequest true "LoginRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.LoginResponse
// @Failure 200 {object} dto.LoginResponse
// @Router /auth/login [post]
func (c *Controller) login(ctx *fiber.Ctx) error {
	var (
		req dto.LoginRequest
		res dto.LoginResponse
	)

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("login user with payload: %s", req)

	res, err = c.Interfaces.AuthViewService.Login(req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Auth
// @Summary Edit user
// @Description Put all mandatory parameter
// @Param Authorization header string true "Authorization"
// @Param EditUserRequest body dto.EditUserRequest true "EditUserRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.EditUserResponse
// @Failure 200 {object} dto.EditUserResponse
// @Router /auth/edit [put]
func (c *Controller) edit(ctx *fiber.Ctx) error {
	var (
		req dto.EditUserRequest
		res dto.EditUserResponse
	)

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("edit user with payload: %s", req)

	user := common.CreateUserContext(ctx, c.Application.AuthService)

	res, err = c.Interfaces.AuthViewService.EditUser(req, user)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Auth
// @Summary Forgot password
// @Description Put all mandatory parameter
// @Param ForgotPasswordRequest body dto.ForgotPasswordRequest true "ForgotPasswordRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.ForgotPasswordRequest
// @Router /auth/forgot-password [post]
func (c *Controller) forgotPassword(ctx *fiber.Ctx) error {
	var (
		req dto.ForgotPasswordRequest
	)

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	c.Shared.Logger.Infof("forgot password request for email: %s", req.Email)

	err = c.Interfaces.AuthViewService.ForgotPassword(req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, nil)
}

// All godoc
// @Tags Auth
// @Summary Reset password
// @Description Put all mandatory parameter
// @Param ResetPasswordRequest body dto.ResetPasswordRequest true "ResetPasswordRequest"
// @Accept  json
// @Produce  json
// @Success 200 {object} dto.ResetPasswordRequest
// @Router /auth/reset-password [post]
func (c *Controller) resetPassword(ctx *fiber.Ctx) error {
	var (
		req dto.ResetPasswordRequest
	)

	err := common.DoCommonRequest(ctx, &req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	err = c.Interfaces.AuthViewService.ResetPassword(req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, nil)
}

// All godoc
// @Tags Auth
// @Summary Get user credential
// @Param Authorization header string true "Authorization"
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200
// @Router /auth/credential [get]
func (c *Controller) userCredential(ctx *fiber.Ctx) error {
	user := common.CreateUserContext(ctx, c.Application.AuthService)

	user = c.Interfaces.AuthViewService.GetUserCredential(user)

	return common.DoCommonSuccessResponse(ctx, user)
}

func NewController(interfaces interfaces.Holder, shared shared.Holder, application application.Holder) Controller {
	return Controller{
		Interfaces:  interfaces,
		Shared:      shared,
		Application: application,
	}
}