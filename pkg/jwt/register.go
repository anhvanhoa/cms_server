package pkgjwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RegisterClaims struct {
	Code string
	Id   string
	jwt.RegisteredClaims
}

func NewRegisterClaims(id, code string, exp time.Time) RegisterClaims {
	return RegisterClaims{
		Code: code,
		Id:   id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			Subject:   "Đăng ký tài khoản",
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}
