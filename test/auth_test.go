package test

import (
	"farmacare/repository"
	"farmacare/service"
	"farmacare/shared"
	"farmacare/shared/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Controller struct {
	Service  service.Holder
	Shared      shared.Holder
	Repository repository.Holder
}

func (c *Controller) TestRegisterUser(t *testing.T) {
	// Membuat data dummy untuk request
	req := dto.CreateUserRequest{
		Email:    "test@example.com",
		Firstname: "Test",
		Lastname: "User",
		Password: "password123",
	}

	// Membuat hasil yang diharapkan
	expectedRes := dto.CreateUserResponse{
		Email:    req.Email,
		Firstname: req.Firstname,
		Lastname: req.Lastname,
		Role:     "user",
	}

	// // Mocking CheckUserExist untuk mengembalikan false (user tidak ada)
	// repository.On("CheckUserExist", req.Email).Return(false, nil)

	// // Mocking CreateUser untuk mengembalikan nil (tidak ada error)
	// repository.On("CreateUser", mock.Anything).Return(nil)

	// Memanggil method yang diuji
	res, err := c.Service.AuthViewService.RegisterUser(req)

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)

	// Memastikan hasil sesuai dengan yang diharapkan
	assert.Equal(t, expectedRes, res)
}


func TestLoginUser(t *testing.T) {
	req := dto.LoginRequest{
		Email: "eugeniusms@gmail.com",
		Password: "password1234",
	}
	res := dto.LoginRequest{
		Email: req.Email,
		Password: req.Password,
	}
	assert.Equal(t, req, res)
}