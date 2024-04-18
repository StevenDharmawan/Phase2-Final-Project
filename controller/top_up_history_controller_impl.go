package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"phase2-final-project/helper"
	"phase2-final-project/model/web/response"
	"phase2-final-project/service"
)

type TopUpHistoryControllerImpl struct {
	service.TopUpHistoryService
}

func NewTopUpHistoryController(topUpHistoryService service.TopUpHistoryService) *TopUpHistoryControllerImpl {
	return &TopUpHistoryControllerImpl{TopUpHistoryService: topUpHistoryService}
}

func (controller *TopUpHistoryControllerImpl) GetAllHistoryById(c *gin.Context) {
	claims := helper.GetClaimsFromCookie(c)
	walletID := claims.WalletID
	responses, errResponse := controller.TopUpHistoryService.GetAllHistoryById(walletID)
	if errResponse != nil {
		c.JSON(errResponse.Code, *errResponse)
		return
	}
	c.JSON(http.StatusOK, response.Response{
		Message: "Success Get All Histories",
		Data:    responses,
	})
}

func (controller *TopUpHistoryControllerImpl) GetAllHistory(c *gin.Context) {
	responses, errResponse := controller.TopUpHistoryService.GetAllHistory()
	if errResponse != nil {
		c.JSON(errResponse.Code, *errResponse)
		return
	}
	c.JSON(http.StatusOK, response.Response{
		Message: "Success Get All Histories",
		Data:    responses,
	})
}
