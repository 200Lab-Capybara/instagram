package userusecase

import (
	userinterface "github.com/nghiatrann0502/instagram-clone/app/internals/services/user/interface"
	"github.com/nghiatrann0502/instagram-clone/components/hasher"
)

type useCase struct {
	userRepository userinterface.UserRepository
	hasher         hasher.Hasher
}

func NewUserUseCase(userRepo userinterface.UserRepository, hasher hasher.Hasher) *useCase {
	return &useCase{
		userRepository: userRepo,
		hasher:         hasher,
	}
}
