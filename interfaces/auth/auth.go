package auth

import (
	"errors"
	"farmacare/application"
	"farmacare/shared"
	"farmacare/shared/common"
	"farmacare/shared/dto"
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
		EditUser(req dto.EditUserRequest, user dto.User) (dto.EditUserResponse, error)
		ForgotPassword(req dto.ForgotPasswordRequest) error
		ResetPassword(req dto.ResetPasswordRequest) error
		GetUserCredential(user dto.User) dto.User
		GoogleLogin(req dto.GoogleLoginRequest) (dto.LoginResponse, error)
	}

	viewService struct {
		application application.Holder
		shared      shared.Holder
	}
)

func (v *viewService) RegisterUser(req dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	var (
		res dto.CreateUserResponse
	)

	isUserExist, _ := v.application.AuthService.CheckUserExist(req.Email)
	if isUserExist {
		return res, errors.New("user already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return res, err
	}

	err = v.application.AuthService.CreateUser(req.TransformToUserModel(string(hashedPassword)))
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

	isUserExist, _ := v.application.AuthService.CheckUserExist(req.Email)
	if isUserExist {
		return res, errors.New("doctor already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return res, err
	}

	err = v.application.AuthService.CreateDoctor(req.TransformToUserModel(string(hashedPassword)))
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

	isUserExist, _ := v.application.AuthService.CheckUserExist(req.Email)
	if isUserExist {
		return res, errors.New("pharmacist already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return res, err
	}

	err = v.application.AuthService.CreatePharmacist(req.TransformToUserModel(string(hashedPassword)))
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

	isUserExist, user := v.application.AuthService.CheckUserExist(req.Email)
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

	isUserExist, user := v.application.AuthService.CheckUserExist(googleData.Email)
	if !isUserExist {
		user = googleData.ToUser()
		err = v.application.AuthService.CreateUser(user)
		if err != nil {
			return res, err
		}
	}

	token, err := common.GenerateToken(v.shared.Env.SecretKey, jwt.MapClaims{
		"id":  user.ID,
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

func (v *viewService) EditUser(req dto.EditUserRequest, user dto.User) (dto.EditUserResponse, error) {
	var (
		res dto.EditUserResponse
	)

	isUserExist, user := v.application.AuthService.CheckUserExist(user.Email)
	if !isUserExist {
		return res, errors.New("no user found for given email")
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

	err := v.application.AuthService.EditUser(user)
	if err != nil {
		return res, err
	}

	res = dto.EditUserResponse{
		Email:    user.Email,
		Fullname: req.Fullname,
	}

	return res, nil
}

func (v *viewService) ForgotPassword(req dto.ForgotPasswordRequest) error {
	isUserExist, user := v.application.AuthService.CheckUserExist(req.Email)
	if !isUserExist {
		return errors.New("no user found for given email")
	}

	v.application.AuthService.RemovePreviousPasswordResetToken(user.ID)

	token := gouid.String(6, gouid.LowerCaseAlphaNum)

	fpwEntry := dto.PasswordReset{
		UserID: user.ID,
		Token:  token,
		Valid:  carbon.Now().AddMinutes(5).ToStdTime(),
	}

	err := v.application.AuthService.CreatePasswordReset(fpwEntry)
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

	err := v.application.AuthService.GetResetToken(req.Token, &pw)
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

	err = v.application.AuthService.EditUser(pw.User)
	if err != nil {
		return err
	}

	go v.application.AuthService.RemovePreviousPasswordResetToken(pw.UserID)

	return nil
}

func (v *viewService) GetUserCredential(user dto.User) dto.User {
	user.HashedPassword = ""
	return user
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}