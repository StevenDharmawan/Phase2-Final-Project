package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"phase2-final-project/helper"
	"phase2-final-project/model/web/request"
	"phase2-final-project/model/web/response"
	"phase2-final-project/service"
)

type RoomControllerImpl struct {
	service.RoomService
}

func NewRoomController(roomService service.RoomService) *RoomControllerImpl {
	return &RoomControllerImpl{RoomService: roomService}
}

func (controller *RoomControllerImpl) GetAllAvailableRooms(c *gin.Context) {
	responses, errResponse := controller.RoomService.GetAllAvailableRooms()
	if errResponse != nil {
		c.JSON(errResponse.Code, *errResponse)
		return
	}
	c.JSON(http.StatusOK, response.Response{
		Message: "Success Get All Rooms",
		Data:    responses,
	})
}

func (controller *RoomControllerImpl) GetAllRooms(c *gin.Context) {
	responses, errResponse := controller.RoomService.GetAllRooms()
	if errResponse != nil {
		c.JSON(errResponse.Code, *errResponse)
		return
	}
	c.JSON(http.StatusOK, response.Response{
		Message: "Success Get All Rooms",
		Data:    responses,
	})
}

func (controller *RoomControllerImpl) CreateRoom(c *gin.Context) {
	roomRequest := request.RoomRequest{}
	err := c.ShouldBindJSON(&roomRequest)
	if err != nil {
		errResponse := helper.ErrBadRequest(err.Error())
		c.JSON(errResponse.Code, errResponse)
		return
	}
	roomResponse, errResponse := controller.RoomService.CreateRoom(roomRequest)
	if errResponse != nil {
		c.JSON(errResponse.Code, *errResponse)
		return
	}
	c.JSON(http.StatusCreated, response.Response{
		Message: "Create Room Success!",
		Data:    roomResponse,
	})
}
