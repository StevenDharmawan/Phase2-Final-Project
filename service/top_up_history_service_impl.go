package service

import (
	"phase2-final-project/helper"
	"phase2-final-project/model/domain"
	"phase2-final-project/model/web/response"
	"phase2-final-project/repository"
)

type TopUpHistoryServiceImpl struct {
	repository.TopUpHistoryRepository
}

func NewTopUpHistoryService(topUpHistoryRepository repository.TopUpHistoryRepository) *TopUpHistoryServiceImpl {
	return &TopUpHistoryServiceImpl{TopUpHistoryRepository: topUpHistoryRepository}
}

func (service *TopUpHistoryServiceImpl) GetAllHistoryById(walletId int) ([]domain.TopUpHistory, *response.ErrorResponse) {
	histories, err := service.TopUpHistoryRepository.GetAllHistoryById(walletId)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	return histories, nil
}

func (service *TopUpHistoryServiceImpl) GetAllHistory() ([]domain.TopUpHistory, *response.ErrorResponse) {
	histories, err := service.TopUpHistoryRepository.GetAllHistory()
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	return histories, nil
}
