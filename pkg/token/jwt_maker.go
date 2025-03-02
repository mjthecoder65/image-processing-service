package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	return &JWTMaker{
		secretKey: secretKey,
	}, nil
}

func (maker *JWTMaker) CreateToken(userID pgtype.UUID, duration time.Duration) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte(maker.secretKey))

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func (maker *JWTMaker) VerifyToken(token string) (*Claims, error) {
	return nil, nil
}
