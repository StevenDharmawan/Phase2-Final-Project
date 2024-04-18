package controller

import "github.com/gin-gonic/gin"

type WalletController interface {
	TopUpWallet(c *gin.Context)
}
