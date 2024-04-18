package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	RegisterAdmin(c *gin.Context)
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	GetUser(c *gin.Context)
}
