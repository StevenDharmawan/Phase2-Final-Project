package repository

import "phase2-final-project/model/domain"

type TopUpHistoryRepository interface {
	GetAllHistoryById(userID int) ([]domain.TopUpHistory, error)
	GetAllHistory() ([]domain.TopUpHistory, error)
	CreateTopUpHistory(request domain.TopUpHistory) error
}
