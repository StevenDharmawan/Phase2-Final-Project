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

type WalletControllerImpl struct {
	service.WalletService
}

func NewWalletController(walletService service.WalletService) *WalletControllerImpl {
	return &WalletControllerImpl{WalletService: walletService}
}

// TopUpWallet
// @Summary Top up wallet balance
// @Description Top up wallet balance
// @Tags wallets
// @Accept json
// @Produce json
// @Param walletRequest body request.WalletRequest true "Wallet top-up input"
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{message=string, data=domain.Wallet} "Successfully topped up wallet balance"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/wallets [post]
func (controller *WalletControllerImpl) TopUpWallet(c *gin.Context) {
	claims := helper.GetClaimsFromCookie(c)
	userID := claims.UserID
	fmt.Println(claims.UserID)
	walletRequest := request.WalletRequest{}
	err := c.ShouldBindJSON(&walletRequest)
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		c.JSON(errResponse.Code, errResponse)
		return
	}
	walletResponse, errResponse := controller.WalletService.TopUpWallet(walletRequest, userID)
	if errResponse != nil {
		c.JSON(errResponse.Code, *errResponse)
		return
	}
	c.JSON(http.StatusOK, response.Response{
		Message: "TopUp Success!",
		Data:    walletResponse,
	})
}
