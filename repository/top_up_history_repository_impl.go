package repository

import (
	"gorm.io/gorm"
	"phase2-final-project/model/domain"
)

type TopUpHistoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewTopUpHistoryRepository(DB *gorm.DB) *TopUpHistoryRepositoryImpl {
	return &TopUpHistoryRepositoryImpl{DB: DB}
}

func (repository *TopUpHistoryRepositoryImpl) GetAllHistoryById(walletId int) ([]domain.TopUpHistory, error) {
	var histories []domain.TopUpHistory
	err := repository.DB.Find(&histories, "wallet_id = ?", walletId).Error
	return histories, err
}

func (repository *TopUpHistoryRepositoryImpl) GetAllHistory() ([]domain.TopUpHistory, error) {
	var histories []domain.TopUpHistory
	err := repository.DB.Find(&histories).Error
	return histories, err
}

func (repository *TopUpHistoryRepositoryImpl) CreateTopUpHistory(request domain.TopUpHistory) error {
	err := repository.DB.Create(&request).Error
	return err
}
