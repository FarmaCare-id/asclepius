package auth

import (
	"farmacare/repository"
	"farmacare/shared"
	"farmacare/shared/common"
	"farmacare/shared/dto"
	"farmacare/shared/exception"
	"farmacare/shared/models"
	"io/ioutil"
	"net/http"

	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt"
	"github.com/golang-module/carbon"
	"github.com/twharmon/gouid"
	"golang.org/x/crypto/bcrypt"
)

type (
	ViewService interface {
		RegisterUser(req dto.CreateUserRequest) (dto.CreateUserResponse, error)
		RegisterDoctor(req dto.CreateDoctorRequest) (dto.CreateDoctorResponse, error)
		RegisterPharmacist(req dto.CreatePharmacistRequest) (dto.CreatePharmacistResponse, error)
		EditUser(req dto.EditUserRequest, ctx dto.SessionContext) (dto.EditUserResponse, error)
		Login(req dto.LoginRequest) (dto.LoginResponse, error)
		ForgotPassword(req dto.ForgotPasswordRequest) error
		ResetPassword(req dto.ResetPasswordRequest) error
		GetUserCredential(ctx dto.SessionContext) models.User
		GoogleLogin(req dto.GoogleLoginRequest) (dto.LoginResponse, error)
	}

	viewService struct {
		repository repository.Holder
		shared      shared.Holder
	}
)

func (v *viewService) RegisterUser(req dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	var (
		res dto.CreateUserResponse
	)

	isUserExist, _ := v.repository.UserRepository.CheckUserExist(req.Email)
	if isUserExist {
		return res, exception.UserAlreadyExist()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return res, err
	}

	err = v.repository.UserRepository.CreateUser(req.TransformToUserModel(string(hashedPassword)))
	if err != nil {
		return res, err
	}

	res = dto.CreateUserResponse{
		Email:    req.Email,
		Fullname: req.Fullname,
		Role:     "user",
	}

	return res, nil
}

func (v *viewService) RegisterDoctor(req dto.CreateDoctorRequest) (dto.CreateDoctorResponse, error) {
	var (
		res dto.CreateDoctorResponse
	)

	isUserExist, _ := v.repository.UserRepository.CheckUserExist(req.Email)
	if isUserExist {
		return res, exception.UserAlreadyExist()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return res, err
	}

	err = v.repository.UserRepository.CreateDoctor(req.TransformToUserModel(string(hashedPassword)))
	if err != nil {
		return res, err
	}

	res = dto.CreateDoctorResponse{
		Email:    req.Email,
		Fullname: req.Fullname,
		Role:     "doctor",
		NoSip:    req.NoSip,
	}

	return res, nil
}

func (v *viewService) RegisterPharmacist(req dto.CreatePharmacistRequest) (dto.CreatePharmacistResponse, error) {
	var (
		res dto.CreatePharmacistResponse
	)

	isUserExist, _ := v.repository.UserRepository.CheckUserExist(req.Email)
	if isUserExist {
		return res, exception.UserAlreadyExist()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return res, err
	}

	err = v.repository.UserRepository.CreatePharmacist(req.TransformToUserModel(string(hashedPassword)))
	if err != nil {
		return res, err
	}

	res = dto.CreatePharmacistResponse{
		Email:    req.Email,
		Fullname: req.Fullname,
		Role:     "pharmacist",
		NoSipa:    req.NoSipa,
	}

	return res, nil
}

func (v *viewService) EditUser(req dto.EditUserRequest, ctx dto.SessionContext) (dto.EditUserResponse, error) {
	var (
		res dto.EditUserResponse
	)

	isUserExist, user := v.repository.UserRepository.CheckUserExist(ctx.User.Email)
	if !isUserExist {
		return res, exception.UserNotFound()
	}

	if req.Fullname != "" {
		user.Fullname = req.Fullname
	}

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
		if err != nil {
			return res, err
		}
		user.HashedPassword = string(hashedPassword)
	}

	if req.Weight != 0 {
		user.Weight = req.Weight
	}

	if req.Height != 0 {
		user.Height = req.Height
	}

	if req.Age != 0 {
		user.Age = req.Age
	}

	if req.NoSip != "" {
		user.NoSip = req.NoSip
	}

	if req.NoSipa != "" {
		user.NoSipa = req.NoSipa
	}

	if req.Specialist != "" {
		user.Specialist = req.Specialist
	}

	if req.Title != "" {
		user.Title = req.Title
	}

	err := v.repository.UserRepository.EditUser(user)
	if err != nil {
		return res, err
	}

	res = dto.EditUserResponse{
		Fullname: req.Fullname,
		Weight:   req.Weight,
		Height:   req.Height,
		Age:      req.Age,
		NoSip:    req.NoSip,
		NoSipa:   req.NoSipa,
		Specialist: req.Specialist,
		Title: req.Title,
	}

	return res, nil
}

func (v *viewService) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	var (
		res dto.LoginResponse
	)

	isUserExist, user := v.repository.UserRepository.CheckUserExist(req.Email)
	if !isUserExist {
		return res, exception.UserNotFound()
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(user.HashedPassword),
		[]byte(req.Password),
	)
	if err != nil {
		return res, exception.IncorrectPassword()
	}

	token, err := common.GenerateToken(v.shared.Env.SecretKey, jwt.MapClaims{
		"id":  user.ID,
		"role": user.Role,
		"exp": carbon.Now().AddDay().Timestamp(),
	})
	if err != nil {
		return res, err
	}

	res = dto.LoginResponse{
		Token: token,
	}

	return res, nil
}

func (v *viewService) GoogleLogin(req dto.GoogleLoginRequest) (dto.LoginResponse, error) {
	var (
		res        dto.LoginResponse
		googleData dto.GoogleData
		googleURL  = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	)

	resp, err := http.Get(googleURL + req.Token)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	json.Unmarshal(response, &googleData)

	isUserExist, user := v.repository.UserRepository.CheckUserExist(googleData.Email)
	if !isUserExist {
		user = googleData.ToUser()
		err = v.repository.UserRepository.CreateUser(user)
		if err != nil {
			return res, err
		}
	}

	token, err := common.GenerateToken(v.shared.Env.SecretKey, jwt.MapClaims{
		"id":  user.ID,
		"role": user.Role,
		"exp": carbon.Now().AddDay().Timestamp(),
	})
	if err != nil {
		return res, err
	}

	res = dto.LoginResponse{
		Token: token,
	}

	return res, nil
}

func (v *viewService) ForgotPassword(req dto.ForgotPasswordRequest) error {
	isUserExist, user := v.repository.UserRepository.CheckUserExist(req.Email)
	if !isUserExist {
		return exception.UserNotFound()
	}

	v.repository.PasswordResetRepository.RemovePreviousPasswordResetToken(user.ID)

	token := gouid.String(6, gouid.LowerCaseAlphaNum)

	fpwEntry := models.PasswordReset{
		UserID: user.ID,
		Token:  token,
		Valid:  carbon.Now().AddMinutes(5).ToStdTime(),
	}

	err := v.repository.PasswordResetRepository.CreatePasswordReset(fpwEntry)
	if err != nil {
		return err
	}

	mail := common.MailerRequest{
		Email: req.Email,
		Name:  user.Fullname,
	}

	go mail.Mailer(v.shared.Env, v.shared.Logger, token)

	return nil
}

func (v *viewService) ResetPassword(req dto.ResetPasswordRequest) error {
	var (
		pw models.PasswordReset
	)

	err := v.repository.PasswordResetRepository.GetResetToken(req.Token, &pw)
	if err != nil {
		return exception.InvalidToken()
	}

	if carbon.Now().ToStdTime().After(pw.Valid) {
		return exception.ExpiredToken()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return err
	}

	pw.User.HashedPassword = string(hashedPassword)

	err = v.repository.UserRepository.EditUser(pw.User)
	if err != nil {
		return err
	}

	go v.repository.PasswordResetRepository.RemovePreviousPasswordResetToken(pw.UserID)

	return nil
}

func (v *viewService) GetUserCredential(ctx dto.SessionContext) models.User {
	ctx.User.HashedPassword = ""
	return ctx.User
}

func NewViewService(repository repository.Holder, shared shared.Holder) ViewService {
	return &viewService{
		repository: repository,
		shared:      shared,
	}
}