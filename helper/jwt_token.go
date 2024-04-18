package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"phase2-final-project/config"
)

func GetToken(tokenString string) (*jwt.Token, error) {
	claims := &config.JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return config.GetJWTKey(), nil
	})
	return token, err
}

func GetClaimsFromCookie(c *gin.Context) *config.JWTClaim {
	tokenString, _ := c.Cookie("token")
	token, _ := GetToken(tokenString)
	claims := token.Claims.(*config.JWTClaim)
	return claims
}
