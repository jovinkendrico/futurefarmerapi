package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("afhaifai12412ajfajf0")

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
