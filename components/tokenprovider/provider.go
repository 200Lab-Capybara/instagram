package tokenprovider

import (
	"github.com/google/uuid"
	"time"
)

type TokenPayload interface {
	GetUserID() uuid.UUID
}

type token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type Provider interface {
	GenerateToken(payload TokenPayload, expiry int) (*token, error)
	ValidateToken(token string) (TokenPayload, error)
}
