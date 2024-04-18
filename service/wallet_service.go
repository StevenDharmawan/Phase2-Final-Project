package service

import (
	"phase2-final-project/model/domain"
	"phase2-final-project/model/web/request"
	"phase2-final-project/model/web/response"
)

type WalletService interface {
	TopUpWallet(request request.WalletRequest, userID int) (*domain.Wallet, *response.ErrorResponse)
}
