package token

import (
	payload "auth_prtc/auth"
	"time"
)

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, *payload.Payload, error) //creates a new token for a specific username and duration
	VerifyToken(token string) (*payload.Payload, error)                                    // it checks if the token is valid or not
}
