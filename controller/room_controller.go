package controller

import (
	"github.com/gin-gonic/gin"
)

type RoomController interface {
	GetAllAvailableRooms(c *gin.Context)
	GetAllRooms(c *gin.Context)
	CreateRoom(c *gin.Context)
}
