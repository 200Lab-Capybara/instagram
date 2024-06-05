package hasher

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type bcryptHasher struct{}

func NewBcryptHasher() *bcryptHasher {
	return &bcryptHasher{}
}

func (b *bcryptHasher) Hash(password, salt string) (string, error) {
	str := fmt.Sprintf("%s.%s", password, salt)
	hash, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (b *bcryptHasher) Verify(password, salt, hash string) bool {
	return password == hash
}

func (b *bcryptHasher) GenSalt(len int8) (string, error) {
	var temp = make([]byte, len)

	_, err := rand.Read(temp)

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(temp), nil
}
