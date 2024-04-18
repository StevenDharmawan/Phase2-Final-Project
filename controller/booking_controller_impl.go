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

type BookingControllerImpl struct {
	service.BookingService
}

func NewBookingController(bookingService service.BookingService) *BookingControllerImpl {
	return &BookingControllerImpl{BookingService: bookingService}
}

func (controller *BookingControllerImpl) CreateBooking(c *gin.Context) {
	claims := helper.GetClaimsFromCookie(c)
	userID := claims.UserID
	userEmail := claims.Email
	bookingRequest := request.BookingRequest{}
	err := c.ShouldBindJSON(&bookingRequest)
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		c.JSON(errResponse.Code, errResponse)
		return
	}
	roomResponse, invoice, errResponse := controller.BookingService.CreateBooking(bookingRequest, userID)
	if errResponse != nil {
		c.JSON(errResponse.Code, *errResponse)
		return
	}
	c.JSON(http.StatusCreated, response.Response{
		Message: "Booking Room Success!",
		Data:    roomResponse,
		Invoice: invoice,
	})
	err = helper.SendEmail(userEmail, "Booking Invoice", fmt.Sprintf("Invoice Link:\n%s", invoice.InvoiceUrl))
}
