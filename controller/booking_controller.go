package controller

import (
	"github.com/gin-gonic/gin"
)

type BookingController interface {
	CreateBooking(c *gin.Context)
}
