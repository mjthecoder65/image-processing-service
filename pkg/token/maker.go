package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type Claims struct {
	UserID pgtype.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

type Maker interface {
	CreateToken(userID pgtype.UUID, duration time.Duration) (string, error)
	VerifyToken(token string) (*Claims, error)
}
