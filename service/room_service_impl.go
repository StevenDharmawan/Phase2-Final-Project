package service

import (
	"phase2-final-project/helper"
	"phase2-final-project/model/domain"
	"phase2-final-project/model/web/request"
	"phase2-final-project/model/web/response"
	"phase2-final-project/repository"
)

type RoomServiceImpl struct {
	repository.RoomRepository
}

func NewRoomService(roomRepository repository.RoomRepository) *RoomServiceImpl {
	return &RoomServiceImpl{RoomRepository: roomRepository}
}

func (service *RoomServiceImpl) GetAllAvailableRooms() ([]domain.Room, *response.ErrorResponse) {
	rooms, err := service.RoomRepository.GetAllAvailableRooms()
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	return rooms, nil
}

func (service *RoomServiceImpl) GetAllRooms() ([]domain.Room, *response.ErrorResponse) {
	rooms, err := service.RoomRepository.GetAllRooms()
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	return rooms, nil
}

func (service *RoomServiceImpl) CreateRoom(request request.RoomRequest) (*domain.Room, *response.ErrorResponse) {
	room := domain.Room{
		RoomNumber:   request.RoomNumber,
		Price:        request.Price,
		Availability: true,
		Category:     request.Category,
	}
	room, err := service.RoomRepository.CreateRoom(room)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	return &room, nil
}
