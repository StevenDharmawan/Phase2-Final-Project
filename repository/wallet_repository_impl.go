package repository

import (
	"gorm.io/gorm"
	"phase2-final-project/model/domain"
)

type WalletRepositoryImpl struct {
	DB *gorm.DB
}

func NewWalletRepository(DB *gorm.DB) *WalletRepositoryImpl {
	return &WalletRepositoryImpl{DB: DB}
}

func (repository *WalletRepositoryImpl) CreateWallet(wallet domain.Wallet) error {
	err := repository.DB.Create(&wallet).Error
	return err
}

func (repository *WalletRepositoryImpl) UpdateWallet(wallet domain.Wallet) (domain.Wallet, error) {
	err := repository.DB.Save(&wallet).Error
	return wallet, err
}

func (repository *WalletRepositoryImpl) GetWalletById(userID int) (domain.Wallet, error) {
	var walletResponse domain.Wallet
	err := repository.DB.Take(&walletResponse, "user_id = ?", userID).Error
	return walletResponse, err
}

func (repository *WalletRepositoryImpl) UpdateBalance(walletID int, balance float64) error {
	err := repository.DB.Model(&domain.Wallet{}).Where("id = ?", walletID).Update("balance", balance).Error
	return err
}
