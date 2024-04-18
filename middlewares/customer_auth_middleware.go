package middlewares

import (
	"github.com/gin-gonic/gin"
	"phase2-final-project/helper"
)

func CustomerAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			errResponse := helper.ErrUnauthorized(err.Error())
			c.JSON(errResponse.Code, errResponse)
			c.Abort()
			return
		}
		token, err := helper.GetToken(tokenString)

		if err != nil || !token.Valid {
			errResponse := helper.ErrUnauthorized(err.Error())
			c.JSON(errResponse.Code, errResponse)
			c.Abort()
			return
		}
		c.Next()
	}
}
