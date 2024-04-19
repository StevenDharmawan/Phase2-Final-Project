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

// GetAllHistoryById
// @Summary Get all top-up histories by wallet ID
// @Description Retrieve all top-up histories associated with a wallet by wallet ID
// @Tags TopUpHistories
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param walletID path int true "Wallet ID"
// @Success 200 {object} response.Response{message=string, data=[]domain.TopUpHistory} "Successfully retrieved all top-up histories"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/users/topuphistories [get]
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

// GetAllHistory
// @Summary Get all top-up histories
// @Description Retrieve all top-up histories
// @Tags TopUpHistories
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.Response{message=string, data=[]domain.TopUpHistory} "Successfully retrieved all top-up histories"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/topuphistories [get]
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
