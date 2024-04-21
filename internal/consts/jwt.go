package consts

import (
	"github.com/golang-jwt/jwt/v5"
	"gokomodo_test/internal/config"
)

var JWT_KEY = []byte(config.NewViper().GetString("key.jwt"))

type JWTClaim struct {
	Id       int
	Email    string
	FullName string
	RoleId   int
	Address  string
	jwt.RegisteredClaims
}
