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

// GetAllAvailableRooms
// @Summary Get all available rooms
// @Description Retrieve all available rooms
// @Tags rooms
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{message=string, data=[]domain.Room}
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/rooms/availability [get]
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

// GetAllRooms
// @Summary Get all rooms
// @Description Retrieve all rooms
// @Tags rooms
// @Accept json
// @Produce json
// @Success 200 {object} response.Response{message=string, data=[]domain.Room}
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/rooms [get]
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

// CreateRoom
// @Summary Create a new room
// @Description Create a new room
// @Tags rooms
// @Accept json
// @Produce json
// @Param roomRequest body request.RoomRequest true "Room input"
// @Success 201 {object} response.Response{message=string, data=domain.Room}
// @Failure 400 {object} response.ErrorResponse "Bad request"
// @Failure 401 {object} response.ErrorResponse "Unauthorized"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @Router /api/v1/rooms [post]
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
