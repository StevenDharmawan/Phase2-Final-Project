package controller

import (
	"github.com/gin-gonic/gin"
)

type TopUpHistoryController interface {
	GetAllHistoryById(c *gin.Context)
	GetAllHistory(c *gin.Context)
}
