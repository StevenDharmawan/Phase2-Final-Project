package repository

import (
	"gorm.io/gorm"
	"phase2-final-project/model/domain"
)

type BookingRepositoryImpl struct {
	DB *gorm.DB
}

func NewBookingRepository(DB *gorm.DB) *BookingRepositoryImpl {
	return &BookingRepositoryImpl{DB: DB}
}

func (repository *BookingRepositoryImpl) CreateBooking(booking domain.Booking) (domain.Booking, error) {
	err := repository.DB.Create(&booking).Error
	return booking, err
}
