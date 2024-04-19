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

// RegisterAdmin
// @Summary Register admin user
// @Description Register admin user
// @Tags users
// @Accept json
// @Produce json
// @Param registerRequest body request.RegisterRequest true "Registration admin input"
// @Success 201 {object} response.Response{message=string, data=response.UserResponse} "Successfully registered admin user"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/users/register/admin [post]
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

// RegisterUser
// @Summary Register user
// @Description Register user
// @Tags users
// @Accept json
// @Produce json
// @Param registerRequest body request.RegisterRequest true "Registration input"
// @Success 201 {object} response.Response{message=string, data=response.UserResponse} "Successfully registered user"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/users/register [post]
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

// LoginUser
// @Summary Login user
// @Description Login user
// @Tags users
// @Accept json
// @Produce json
// @Param loginRequest body request.LoginRequest true "Login input"
// @Success 200 {object} response.Response{message=string} "Successfully logged in user"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/users/login [post]
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

// UpdateUser
// @Summary Update user profile
// @Description Update user profile
// @Tags users
// @Accept json
// @Produce json
// @Param updateUserRequest body request.UpdateUserRequest true "Update user profile input"
// @Success 200 {object} response.Response{message=string, data=response.UserResponse} "Successfully updated user profile"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/users [put]
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

// GetUser
// @Summary Get user profile
// @Description Get user profile
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{message=string, data=domain.User} "Successfully retrieved user profile"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/users [get]
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
