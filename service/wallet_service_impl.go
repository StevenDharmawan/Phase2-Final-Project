package service

import (
	"fmt"
	"phase2-final-project/helper"
	"phase2-final-project/model/domain"
	"phase2-final-project/model/web/request"
	"phase2-final-project/model/web/response"
	"phase2-final-project/repository"
	"time"
)

type WalletServiceImpl struct {
	repository.WalletRepository
	repository.TopUpHistoryRepository
}

func NewWalletService(walletRepository repository.WalletRepository, topUpHistoryRepository repository.TopUpHistoryRepository) *WalletServiceImpl {
	return &WalletServiceImpl{
		WalletRepository:       walletRepository,
		TopUpHistoryRepository: topUpHistoryRepository,
	}
}

func (service *WalletServiceImpl) TopUpWallet(request request.WalletRequest, userID int) (*domain.Wallet, *response.ErrorResponse) {
	wallet, err := service.WalletRepository.GetWalletById(userID)
	wallet.Balance += request.AmountTopUp
	fmt.Println(wallet)
	wallet, err = service.WalletRepository.UpdateWallet(wallet)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	topUpHistory := domain.TopUpHistory{
		WalletID:      wallet.ID,
		AmountTopUp:   request.AmountTopUp,
		TopUpDateTime: time.Now(),
	}
	err = service.TopUpHistoryRepository.CreateTopUpHistory(topUpHistory)
	if err != nil {
		errResponse := helper.ErrInternalServer(err.Error())
		return nil, &errResponse
	}
	return &wallet, nil
}
