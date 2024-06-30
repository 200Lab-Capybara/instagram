package common

import (
	"github.com/google/uuid"
)

const (
	AccessTokenExpireDuration  = 60 * 60
	RefreshTokenExpireDuration = 60 * 60 * 24 * 7
)

type TokenPayload struct {
	Id uuid.UUID `json:"id"`
}

func (t *TokenPayload) GetUserID() uuid.UUID {
	return t.Id
}
