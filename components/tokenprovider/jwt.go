package tokenprovider

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"instagram/common"
	"time"
)

type myClaims struct {
	common.TokenPayload
	jwt.RegisteredClaims
}

type jwtProvider struct {
	secretKey string
}

func NewJWTProvider(secretKey string) Provider {
	return &jwtProvider{secretKey: secretKey}
}

func (j jwtProvider) GenerateToken(payload TokenPayload, expiry int) (*token, error) {
	now := time.Now()

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		common.TokenPayload{
			Id: payload.GetUserID(),
		},
		jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			ExpiresAt: jwt.NewNumericDate(now.Local().Add(time.Second * time.Duration(expiry))),
			IssuedAt:  jwt.NewNumericDate(now.Local()),
		},
	})

	tokenString, err := claims.SignedString([]byte(j.secretKey))
	if err != nil {
		return nil, err
	}

	return &token{
		Token:   tokenString,
		Created: now,
		Expiry:  expiry,
	}, nil
}

func (j jwtProvider) ValidateToken(token string) (TokenPayload, error) {
	claims := &myClaims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return &claims.TokenPayload, nil
}
