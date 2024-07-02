package users

import (
	dtoUsers "backend/dtos/users"
	services "backend/services/contracts"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserControllerContract interface {
	RegisterUser(c *gin.Context)
	Login(c *gin.Context)
}

type UserControllerImplementation struct {
	UserService services.UserServiceContract
}

func (userController UserControllerImplementation) RegisterUser(context *gin.Context) {

	var userRequest dtoUsers.AuthDto
	context.BindJSON(&userRequest)
	response, err := userController.UserService.RegisterUser(userRequest)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}
	context.JSON(http.StatusCreated, response)
}

func (userController UserControllerImplementation) Login(context *gin.Context) {
	var loginRequest dtoUsers.AuthDto
	context.BindJSON(&loginRequest)
	response, err := userController.UserService.Login(loginRequest)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}
	context.JSON(http.StatusOK, response)
}
