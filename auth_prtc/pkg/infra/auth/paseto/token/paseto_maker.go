package token

import (
	ap "auth_prtc/pkg/infra/auth"
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

// Paseto maker
type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// creating a new pasetoMaker
func NewPasetoMaker(something string) (ap.Maker, error) {
	//checking the size, ps: const size is 32
	if len(something) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("wrong size for key")
	}

	//
	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(something),
	}

	return maker, nil
}

// creating a new token based on a specific username duration
func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, *ap.Payload, error) {
	// declaring a payload
	payload, err := ap.NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}
	// encrypting the payload and declaring the token
	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)

	// return the token
	return token, payload, err
}

// verifying the token for validity
func (maker *PasetoMaker) VerifyToken(token string) (*ap.Payload, error) {
	//decrypt the payload (token included)
	payload := &ap.Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ap.ErrInvalidToken
	}

	//check for the validity
	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	//return the patload
	return payload, nil
}
