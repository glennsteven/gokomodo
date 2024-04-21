package consts

import (
	"github.com/golang-jwt/jwt/v5"
	"gokomodo_test/internal/config"
)

var JWT_KEY = []byte(config.NewViper().GetString("key.jwt"))

type JWTClaim struct {
	Email string
	jwt.RegisteredClaims
}
