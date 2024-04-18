package service

import (
	"phase2-final-project/helper"
	"phase2-final-project/model/domain"
	"phase2-final-project/model/web/request"
	"phase2-final-project/model/web/response"
	"phase2-final-project/repository"
	"time"
)

type BookingServiceImpl struct {
	repository.BookingRepository
	repository.RoomRepository
	repository.WalletRepository
	repository.UserRepository
}

func NewBookingService(bookingRepository repository.BookingRepository, roomRepository repository.RoomRepository, walletRepository repository.WalletRepository, userRepository repository.UserRepository) *BookingServiceImpl {
	return &BookingServiceImpl{
		BookingRepository: bookingRepository,
		RoomRepository:    roomRepository,
		WalletRepository:  walletRepository,
		UserRepository:    userRepository}
}

func (service *BookingServiceImpl) CreateBooking(request request.BookingRequest, userID int) (*domain.Booking, *domain.Invoice, *response.ErrorResponse) {
	room, err := service.RoomRepository.GetRoomById(request.RoomID)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, nil, &errResponse
	}
	if room.Availability == false {
		errResponse := helper.ErrBadRequest("Room Sedang Tidak Tersedia")
		return nil, nil, &errResponse
	}
	currentTime := time.Now()
	checkinDate, _ := time.Parse("2006-01-02", currentTime.Format("2006-01-02"))
	checkoutDate, _ := time.Parse("2006-01-02", request.CheckoutDate)

	days := checkoutDate.Sub(checkinDate).Hours() / 24
	totalPrice := days * room.Price
	wallet, err := service.WalletRepository.GetWalletById(userID)
	wallet.Balance -= totalPrice
	if wallet.Balance < 0 {
		errResponse := helper.ErrBadRequest("Saldo tidak mencukupi untuk melakukan transaksi ini")
		return nil, nil, &errResponse
	}

	err = service.WalletRepository.UpdateBalance(wallet.ID, wallet.Balance)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, nil, &errResponse
	}
	booking := domain.Booking{
		UserID:       userID,
		RoomID:       request.RoomID,
		CheckinDate:  checkinDate,
		CheckoutDate: checkoutDate,
		TotalPrice:   totalPrice,
	}
	booking, err = service.BookingRepository.CreateBooking(booking)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, nil, &errResponse
	}
	err = service.RoomRepository.UpdateRoom(request.RoomID)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, nil, &errResponse
	}
	user, err := service.UserRepository.GetUserById(userID)
	invoiceRes, err := helper.CreateInvoice(user, totalPrice)
	if err != nil {
		errResponse := helper.ErrInternalServer("Error while creating invoices")
		return nil, nil, &errResponse
	}
	return &booking, invoiceRes, nil
}
