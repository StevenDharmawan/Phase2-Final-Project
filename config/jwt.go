package config

import "github.com/golang-jwt/jwt/v5"

type JWTClaim struct {
	UserID   int
	WalletID int
	Email    string
	Role     string
	jwt.RegisteredClaims
}

func GetJWTKey() []byte {
	return []byte(GetEnv("JWT_KEY"))
}
