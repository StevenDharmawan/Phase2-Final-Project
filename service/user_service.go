package service

import (
	"phase2-final-project/model/web/request"
	"phase2-final-project/model/web/response"
)

type UserService interface {
	RegisterAdmin(request request.RegisterRequest) (*response.UserResponse, *response.ErrorResponse)
	RegisterUser(request request.RegisterRequest) (*response.UserResponse, *response.ErrorResponse)
	LoginUser(loginRequest request.LoginRequest) (*string, *response.ErrorResponse)
	UpdateUser(request request.UpdateUserRequest, userID int) (*response.UserResponse, *response.ErrorResponse)
	GetUser(userID int) (*response.UserResponse, *response.ErrorResponse)
}
