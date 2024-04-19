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

// CreateBooking
// @Summary Create a new booking
// @Description Create a new booking for a room
// @Tags bookings
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer token"
// @Param bookingRequest body request.BookingRequest true "Booking input"
// @Success 201 {object} response.Response{message=string, data=domain.Booking, invoice=domain.Invoice} "Booking Room Success!"
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/bookings [post]
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
