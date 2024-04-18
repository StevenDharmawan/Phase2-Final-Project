package service

import (
	"phase2-final-project/model/domain"
	"phase2-final-project/model/web/response"
)

type TopUpHistoryService interface {
	GetAllHistoryById(userID int) ([]domain.TopUpHistory, *response.ErrorResponse)
	GetAllHistory() ([]domain.TopUpHistory, *response.ErrorResponse)
}
