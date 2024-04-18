package helper

import (
	"phase2-final-project/model/domain"
	"phase2-final-project/model/web/response"
)

func ToUserResponse(user domain.User) response.UserResponse {
	return response.UserResponse{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Age:      user.Age,
		Wallet:   user.Wallet,
	}
}
