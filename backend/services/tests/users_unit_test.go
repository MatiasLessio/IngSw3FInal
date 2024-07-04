package test

import (
	dto "backend/dtos/users"
	e "backend/utilities"
	"testing"

	"github.com/stretchr/testify/assert"
)

type UserMockClient struct{}

func NewUserMockClient() *UserMockClient {
	return &UserMockClient{}
}

type userService struct {
	userClient *UserMockClient
}

func (s *userService) Login(request dto.AuthDto) (dto.LoginResponseDto, e.ApiError) {
	if request.Username == "admin" && request.Password == "123" {
		return dto.LoginResponseDto{
			Username: request.Username,
			Token:    "tokenTestResponse",
		}, nil
	} else {
		return dto.LoginResponseDto{}, e.NewBadRequestApiError("wrong username or password")
	}
}
func (s *userService) Register(request dto.AuthDto) (dto.AuthDto, e.ApiError) {
	return request, nil
}

func TestLogin(t *testing.T) {

	service := &userService{userClient: NewUserMockClient()}
	request := dto.AuthDto{
		Username: "admin",
		Password: "123",
	}

	_, err := service.Login(request)
	assert.NoError(t, err, "wrong pw or username")
}

func TestRegister(t *testing.T) {

	service := &userService{userClient: NewUserMockClient()}
	request := dto.AuthDto{
		Username: "admin",
		Password: "123",
	}

	_, err := service.Register(request)
	assert.NoError(t, err, "error while creating a new account")
}
