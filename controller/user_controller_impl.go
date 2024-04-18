package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"phase2-final-project/helper"
	"phase2-final-project/model/web/request"
	"phase2-final-project/model/web/response"
	"phase2-final-project/service"
)

type UserControllerImpl struct {
	service.UserService
}

func NewUserController(userService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{UserService: userService}
}

func (controller *UserControllerImpl) RegisterAdmin(c *gin.Context) {
	registerRequest := request.RegisterRequest{}
	err := c.ShouldBindJSON(&registerRequest)
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		c.JSON(errResponse.Code, errResponse)
		return
	}
	registerResponse, errResponse := controller.UserService.RegisterAdmin(registerRequest)
	fmt.Println(errResponse)
	if errResponse != nil {
		c.JSON(errResponse.Code, *errResponse)
		return
	}
	c.JSON(http.StatusCreated, response.Response{
		Message: "Register Success!",
		Data:    registerResponse,
	})
}

func (controller *UserControllerImpl) RegisterUser(c *gin.Context) {
	registerRequest := request.RegisterRequest{}
	err := c.ShouldBindJSON(&registerRequest)
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		c.JSON(errResponse.Code, errResponse)
		return
	}
	registerResponse, errResponse := controller.UserService.RegisterUser(registerRequest)
	fmt.Println(errResponse)
	if errResponse != nil {
		c.JSON(errResponse.Code, *errResponse)
		return
	}
	c.JSON(http.StatusCreated, response.Response{
		Message: "Register Success!",
		Data:    registerResponse,
	})
	err = helper.SendEmail(registerResponse.Email, "Registration Success!", "Your account has been successfully created")
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		c.JSON(errResponse.Code, errResponse)
		return
	}
}

func (controller *UserControllerImpl) LoginUser(c *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := c.ShouldBindJSON(&loginRequest)
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		c.JSON(errResponse.Code, errResponse)
		return
	}
	token, errResponse := controller.UserService.LoginUser(loginRequest)
	if errResponse != nil {
		c.JSON(errResponse.Code, *errResponse)
		return
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    *token,
		Path:     "/",
		HttpOnly: true,
	})
	c.JSON(http.StatusOK, response.Response{
		Message: "Login Success!",
	})
}

func (controller *UserControllerImpl) UpdateUser(c *gin.Context) {
	claims := helper.GetClaimsFromCookie(c)
	userID := claims.UserID
	updateUserRequest := request.UpdateUserRequest{}
	err := c.ShouldBindJSON(&updateUserRequest)
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		c.JSON(errResponse.Code, errResponse)
		return
	}
	registerResponse, errResponse := controller.UserService.UpdateUser(updateUserRequest, userID)
	if errResponse != nil {
		c.JSON(errResponse.Code, *errResponse)
		return
	}
	c.JSON(http.StatusOK, response.Response{
		Message: "Update Success!",
		Data:    registerResponse,
	})
}

func (controller *UserControllerImpl) GetUser(c *gin.Context) {
	claims := helper.GetClaimsFromCookie(c)
	userID := claims.UserID
	userResponse, errResponse := controller.UserService.GetUser(userID)
	if errResponse != nil {
		c.JSON(errResponse.Code, *errResponse)
		return
	}
	c.JSON(http.StatusOK, response.Response{
		Message: "Success Get User",
		Data:    userResponse,
	})
}
