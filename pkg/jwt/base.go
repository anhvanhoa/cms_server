package pkgjwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWT interface {
	GenRegisterToken(data RegisterClaims) (string, error)
	VerifyRegisterToken(token string) (*RegisterClaims, error)
}

type jwtImpl struct {
	secretKey string
}

func NewJWT(secretKey string) JWT {
	return &jwtImpl{
		secretKey: secretKey,
	}
}

func (j *jwtImpl) generateToken(data jwt.Claims) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, data)
}

func (j *jwtImpl) GenRegisterToken(data RegisterClaims) (string, error) {
	token := j.generateToken(data)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtImpl) verifyClaim(token string, data jwt.Claims) (*jwt.Token, error) {
	t, err := jwt.ParseWithClaims(token, data, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKeyType
		}
		return []byte(j.secretKey), nil
	})
	return t, err
}

func (j *jwtImpl) VerifyRegisterToken(token string) (*RegisterClaims, error) {
	t, err := j.verifyClaim(token, &RegisterClaims{})
	if err != nil {
		return nil, err
	}
	claim, ok := t.Claims.(*RegisterClaims)
	if ok {
		return claim, nil
	}
	return nil, ErrParseToken
}
