package repository

import "phase2-final-project/model/domain"

type WalletRepository interface {
	CreateWallet(wallet domain.Wallet) error
	UpdateWallet(wallet domain.Wallet) (domain.Wallet, error)
	GetWalletById(userID int) (domain.Wallet, error)
	UpdateBalance(walletID int, balance float64) error
}
