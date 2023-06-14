package auth

import (
	"farmacare/repository"
	"farmacare/service"
	"farmacare/shared"
	"farmacare/shared/common"
	"farmacare/shared/dto"

	"github.com/gofiber/fiber/v2"
)
type Controller struct {
	Service  service.Holder
	Shared      shared.Holder
	Repository repository.Holder
}

func (c *Controller) Routes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", c.register)
	auth.Post("/register/doctor", c.registerDoctor)
	auth.Post("/register/pharmacist", c.registerPharmacist)
	auth.Post("/login", c.login)
	auth.Post("/login-google", c.loginGoogle)
	auth.Post("/logout", c.Shared.Middleware.AuthMiddleware, c.logout)
	auth.Post("/forgot-password", c.forgotPassword)
	auth.Post("/reset-password", c.resetPassword)
	auth.Get("/credential", c.Shared.Middleware.AuthMiddleware, c.userCredential)
}

// All godoc
// @Tags Auth
// @Summary Register New User
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

	res, err = c.Service.AuthViewService.RegisterUser(req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Auth
// @Summary Register New Doctor
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

	res, err = c.Service.AuthViewService.RegisterDoctor(req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Auth
// @Summary Register New Pharmacist
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

	res, err = c.Service.AuthViewService.RegisterPharmacist(req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Auth
// @Summary Login User
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

	res, err = c.Service.AuthViewService.GoogleLogin(req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Auth
// @Summary Login User Google
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

	res, err = c.Service.AuthViewService.Login(req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, res)
}

// All godoc
// @Tags Auth
// @Summary Logout User
// @Param Authorization header string true "Authorization"
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200
// @Router /auth/logout [post]
func (c *Controller) logout(ctx *fiber.Ctx) error {
	context := common.CreateContext(ctx)

	token, err := c.Service.AuthViewService.Logout(context)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, token)
}

// All godoc
// @Tags Auth
// @Summary Forgot Password
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

	err = c.Service.AuthViewService.ForgotPassword(req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, nil)
}

// All godoc
// @Tags Auth
// @Summary Reset Password
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

	err = c.Service.AuthViewService.ResetPassword(req)
	if err != nil {
		return common.DoCommonErrorResponse(ctx, err)
	}

	return common.DoCommonSuccessResponse(ctx, nil)
}

// All godoc
// @Tags Auth
// @Summary Get User Credential
// @Param Authorization header string true "Authorization"
// @Description Put all mandatory parameter
// @Accept  json
// @Produce  json
// @Success 200
// @Router /auth/credential [get]
func (c *Controller) userCredential(ctx *fiber.Ctx) error {
	context := common.CreateContext(ctx)

	user := c.Service.AuthViewService.GetUserCredential(context)

	return common.DoCommonSuccessResponse(ctx, user)
}

func NewController(service service.Holder, shared shared.Holder, repository repository.Holder) Controller {
	return Controller{
		Service:  service,
		Shared:      shared,
		Repository: repository,
	}
}