package service

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"phase2-final-project/config"
	"phase2-final-project/helper"
	"phase2-final-project/model/domain"
	"phase2-final-project/model/web/request"
	"phase2-final-project/model/web/response"
	"phase2-final-project/repository"
	"time"
)

type UserServiceImpl struct {
	repository.UserRepository
	repository.WalletRepository
}

func NewUserService(userRepository repository.UserRepository, walletRepository repository.WalletRepository) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository:   userRepository,
		WalletRepository: walletRepository,
	}
}

func (service *UserServiceImpl) RegisterAdmin(request request.RegisterRequest) (*response.UserResponse, *response.ErrorResponse) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	user := domain.User{
		FullName: request.FullName,
		Email:    request.Email,
		Password: string(hashPassword),
		Age:      request.Age,
		Role:     "admin",
	}
	user, err := service.UserRepository.CreateUser(user)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	registerResponse := helper.ToUserResponse(user)
	return &registerResponse, nil
}

func (service *UserServiceImpl) RegisterUser(request request.RegisterRequest) (*response.UserResponse, *response.ErrorResponse) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	user := domain.User{
		FullName: request.FullName,
		Email:    request.Email,
		Password: string(hashPassword),
		Age:      request.Age,
	}
	user, err := service.UserRepository.CreateUser(user)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	wallet := domain.Wallet{
		UserID: user.ID,
	}
	err = service.WalletRepository.CreateWallet(wallet)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	registerResponse, err := service.UserRepository.GetUserById(user.ID)
	return &registerResponse, nil
}

func (service *UserServiceImpl) LoginUser(request request.LoginRequest) (*string, *response.ErrorResponse) {
	user, err := service.UserRepository.FindByEmail(request.Email)
	if err != nil {
		errResponse := helper.ErrUnauthorized(err.Error())
		return nil, &errResponse
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		errResponse := helper.ErrUnauthorized(err.Error())
		return nil, &errResponse
	}
	expTime := time.Now().Add(time.Minute * 5)
	claims := &config.JWTClaim{
		UserID:   user.ID,
		WalletID: user.Wallet.ID,
		Email:    user.Email,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(config.GetJWTKey())
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	return &token, nil
}

func (service *UserServiceImpl) UpdateUser(request request.UpdateUserRequest, userID int) (*response.UserResponse, *response.ErrorResponse) {
	if request.Password != "" {
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		request.Password = string(hashPassword)
	}
	user := domain.User{
		ID:       userID,
		FullName: request.FullName,
		Email:    request.Email,
		Password: request.Password,
		Age:      request.Age,
	}
	user, err := service.UserRepository.UpdateUser(user)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	updateResponse := helper.ToUserResponse(user)
	return &updateResponse, nil
}

func (service *UserServiceImpl) GetUser(userID int) (*response.UserResponse, *response.ErrorResponse) {
	user, err := service.UserRepository.GetUserById(userID)
	if err != nil {
		errResponse := helper.ErrNotFound(err.Error())
		return nil, &errResponse
	}
	return &user, nil
}
