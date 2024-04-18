package repository

import (
	"gorm.io/gorm"
	"phase2-final-project/model/domain"
)

type RoomRepositoryImpl struct {
	DB *gorm.DB
}

func NewRoomRepository(db *gorm.DB) *RoomRepositoryImpl {
	return &RoomRepositoryImpl{DB: db}
}

func (repository *RoomRepositoryImpl) GetAllAvailableRooms() ([]domain.Room, error) {
	var rooms []domain.Room
	err := repository.DB.Find(&rooms, "availability = ?", true).Error
	return rooms, err
}

func (repository *RoomRepositoryImpl) GetAllRooms() ([]domain.Room, error) {
	var rooms []domain.Room
	err := repository.DB.Find(&rooms).Error
	return rooms, err
}

func (repository *RoomRepositoryImpl) CreateRoom(room domain.Room) (domain.Room, error) {
	err := repository.DB.Create(&room).Error
	return room, err
}

func (repository *RoomRepositoryImpl) UpdateRoom(roomID int) error {
	err := repository.DB.Model(&domain.Room{}).Where("id = ?", roomID).Update("availability", false).Error
	return err
}

func (repository *RoomRepositoryImpl) GetRoomById(roomID int) (domain.Room, error) {
	var room domain.Room
	err := repository.DB.Find(&room, "id = ?", roomID).Error
	return room, err
}
