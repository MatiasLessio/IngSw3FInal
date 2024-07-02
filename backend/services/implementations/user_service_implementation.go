package services

import (
	client "backend/clients/users"
	dto "backend/dtos/users"
	"backend/models"
	e "backend/utilities"
	"crypto/md5"
	"encoding/hex"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

type UserServiceImplementation struct{}

func NewUserServiceImpl() UserServiceImplementation {
	return UserServiceImplementation{}
}

func (UserServiceImplementation) RegisterUser(request dto.AuthDto) (dto.AuthDto, e.ApiError) {
	var user models.User

	user.Username = request.Username
	user.Password = request.Password

	hash := md5.New()
	hash.Write([]byte(request.Password))
	user.Password = hex.EncodeToString(hash.Sum(nil))

	user, err := client.RegisterUser(user)
	if err != nil {
		return request, e.NewBadRequestApiError("error saving user")
	}
	request.Username = user.Username

	return request, nil
}

func (UserServiceImplementation) Login(request dto.AuthDto) (dto.LoginResponseDto, e.ApiError) {
	var user, err = client.GetUser(request.Username)
	var authResponse dto.LoginResponseDto
	if err != nil {
		return authResponse, e.NewBadRequestApiError("user not found")
	}

	var pswMd5 = md5.Sum([]byte(request.Password))
	pswMd5String := hex.EncodeToString(pswMd5[:])

	if pswMd5String == user.Password {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.UserId,
		})
		tokenString, _ := token.SignedString(jwtKey)
		authResponse.Token = tokenString
		authResponse.Username = user.Username

		return authResponse, nil
	} else {
		return authResponse, e.NewBadRequestApiError("Wrong password")
	}
}
