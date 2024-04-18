package repository

import (
	"gorm.io/gorm"
	"phase2-final-project/model/domain"
	"phase2-final-project/model/web/response"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: DB}
}

func (repository *UserRepositoryImpl) CreateUser(user domain.User) (domain.User, error) {
	err := repository.DB.Create(&user).Error
	return user, err
}

func (repository *UserRepositoryImpl) FindByEmail(userEmail string) (domain.User, error) {
	var user domain.User
	err := repository.DB.Preload("Wallet").Take(&user, "Email = ?", userEmail).Error
	return user, err
}

func (repository *UserRepositoryImpl) UpdateUser(user domain.User) (domain.User, error) {
	err := repository.DB.Save(&user).Error
	return user, err
}

func (repository *UserRepositoryImpl) GetUserById(userID int) (response.UserResponse, error) {
	var userResponse response.UserResponse
	err := repository.DB.Model(&domain.User{}).Preload("Wallet").Take(&userResponse, "ID = ?", userID).Error
	return userResponse, err
}
