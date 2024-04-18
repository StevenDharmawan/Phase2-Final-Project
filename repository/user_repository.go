package repository

import (
	"phase2-final-project/model/domain"
	"phase2-final-project/model/web/response"
)

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	FindByEmail(userEmail string) (domain.User, error)
	UpdateUser(user domain.User) (domain.User, error)
	GetUserById(userID int) (response.UserResponse, error)
}
