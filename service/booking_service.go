package service

import (
	"phase2-final-project/model/domain"
	"phase2-final-project/model/web/request"
	"phase2-final-project/model/web/response"
)

type BookingService interface {
	CreateBooking(request request.BookingRequest, userID int) (*domain.Booking, *domain.Invoice, *response.ErrorResponse)
}
