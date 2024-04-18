package service

import (
	"phase2-final-project/model/domain"
	"phase2-final-project/model/web/request"
	"phase2-final-project/model/web/response"
)

type RoomService interface {
	GetAllAvailableRooms() ([]domain.Room, *response.ErrorResponse)
	GetAllRooms() ([]domain.Room, *response.ErrorResponse)
	CreateRoom(request request.RoomRequest) (*domain.Room, *response.ErrorResponse)
}
