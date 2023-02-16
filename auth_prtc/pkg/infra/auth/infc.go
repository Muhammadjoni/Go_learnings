package auth

import (
	"time"
)

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, *Payload, error) //creates a new token for a specific username and duration
	VerifyToken(token string) (*Payload, error)                                    // it checks if the token is valid or not
}
