package services

import (
	dto "backend/dtos/users"
	e "backend/utilities"
)

type UserServiceContract interface {
	Login(request dto.AuthDto) (dto.LoginResponseDto, e.ApiError)
	RegisterUser(request dto.AuthDto) (dto.AuthDto, e.ApiError)
}
