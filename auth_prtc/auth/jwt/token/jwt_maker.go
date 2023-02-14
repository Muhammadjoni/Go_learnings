package token

import (
	ap "auth_prtc/auth"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSecretKeySize = 32

// JSON Web Token maker
type JWTMaker struct {
	secretKey string
}

// create a new JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: min size %d", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil // Waiting for Maker server logicccccccccccccccccccccccc
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, *ap.Payload, error) {
	payload, err := ap.NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err
}

// verifying the token validity
func (maker *JWTMaker) VerifyToken(token string) (*ap.Payload, error) {
	kFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ap.ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &ap.Payload{}, kFunc)
	if err != nil {
		vrf, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(vrf.Inner, ap.ErrExpiredToken) {
			return nil, ap.ErrExpiredToken
		}
		return nil, ap.ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*ap.Payload)
	if !ok {
		return nil, ap.ErrInvalidToken
	}

	return payload, nil
}
