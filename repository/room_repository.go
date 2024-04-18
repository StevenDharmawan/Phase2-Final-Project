package repository

import "phase2-final-project/model/domain"

type RoomRepository interface {
	GetAllAvailableRooms() ([]domain.Room, error)
	GetAllRooms() ([]domain.Room, error)
	CreateRoom(room domain.Room) (domain.Room, error)
	UpdateRoom(roomID int) error
	GetRoomById(roomId int) (domain.Room, error)
}
