package middlewares

import (
	"github.com/gin-gonic/gin"
	"phase2-final-project/config"
	"phase2-final-project/helper"
)

func AdminAuthMiddleware() gin.HandlerFunc {
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
		claims := token.Claims.(*config.JWTClaim)
		if claims.Role != "admin" {
			errResponse := helper.ErrUnauthorized("Admin Role Required")
			c.JSON(errResponse.Code, errResponse)
			c.Abort()
			return
		}
		c.Next()
	}
}
