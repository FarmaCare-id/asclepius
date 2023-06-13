package auth

import (
	"errors"
	"farmacare/repository"
	"farmacare/shared"
	"farmacare/shared/common"
	"farmacare/shared/dto"
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

	isUserExist, _ := v.repository.AuthService.CheckUserExist(req.Email)
	if isUserExist {
		return res, errors.New("user already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return res, err
	}

	err = v.repository.AuthService.CreateUser(req.TransformToUserModel(string(hashedPassword)))
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

	isUserExist, _ := v.repository.AuthService.CheckUserExist(req.Email)
	if isUserExist {
		return res, errors.New("doctor already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return res, err
	}

	err = v.repository.AuthService.CreateDoctor(req.TransformToUserModel(string(hashedPassword)))
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

	isUserExist, _ := v.repository.AuthService.CheckUserExist(req.Email)
	if isUserExist {
		return res, errors.New("pharmacist already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return res, err
	}

	err = v.repository.AuthService.CreatePharmacist(req.TransformToUserModel(string(hashedPassword)))
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

func (v *viewService) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	var (
		res dto.LoginResponse
	)

	isUserExist, user := v.repository.AuthService.CheckUserExist(req.Email)
	if !isUserExist {
		return res, errors.New("no user found for given email")
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(user.HashedPassword),
		[]byte(req.Password),
	)
	if err != nil {
		return res, errors.New("incorrect password")
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

	isUserExist, user := v.repository.AuthService.CheckUserExist(googleData.Email)
	if !isUserExist {
		user = googleData.ToUser()
		err = v.repository.AuthService.CreateUser(user)
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
	isUserExist, user := v.repository.AuthService.CheckUserExist(req.Email)
	if !isUserExist {
		return errors.New("no user found for given email")
	}

	v.repository.AuthService.RemovePreviousPasswordResetToken(user.ID)

	token := gouid.String(6, gouid.LowerCaseAlphaNum)

	fpwEntry := dto.PasswordReset{
		UserID: user.ID,
		Token:  token,
		Valid:  carbon.Now().AddMinutes(5).ToStdTime(),
	}

	err := v.repository.AuthService.CreatePasswordReset(fpwEntry)
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
		pw dto.PasswordReset
	)

	err := v.repository.AuthService.GetResetToken(req.Token, &pw)
	if err != nil {
		return errors.New("token is invalid")
	}

	if carbon.Now().ToStdTime().After(pw.Valid) {
		return errors.New("token is expired")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return err
	}

	pw.User.HashedPassword = string(hashedPassword)

	err = v.repository.ProfileService.EditUserProfile(pw.User)
	if err != nil {
		return err
	}

	go v.repository.AuthService.RemovePreviousPasswordResetToken(pw.UserID)

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