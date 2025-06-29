package pkgjwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	Id       string
	FullName string
	jwt.RegisteredClaims
}

func NewAuthClaims(id, fullName string, exp time.Time) AuthClaims {
	return AuthClaims{
		Id:       id,
		FullName: fullName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			Subject:   "Xác thực tài khoản " + id,
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}
