package repository

import "phase2-final-project/model/domain"

type BookingRepository interface {
	CreateBooking(booking domain.Booking) (domain.Booking, error)
}
