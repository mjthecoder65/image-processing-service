package token

import "time"

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	return &JWTMaker{
		secretKey: secretKey,
	}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	return "", nil
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	return nil, nil
}
