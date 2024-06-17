package userusecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	usermodel "github.com/nghiatrann0502/instagram-clone/app/internals/services/user/model"
	"github.com/nghiatrann0502/instagram-clone/components/hasher"
)

type registerUseCase struct {
	registerRepository RegisterRepository
	hasher             hasher.Hasher
}

func NewRegisterUseCase(registerRepository RegisterRepository, hasher hasher.Hasher) RegisterUseCase {
	return &registerUseCase{
		registerRepository: registerRepository,
		hasher:             hasher,
	}
}

type RegisterUseCase interface {
	Execute(ctx context.Context, user *usermodel.UserCreation) (*uuid.UUID, error)
}

type RegisterRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*usermodel.User, error)
	CreateNewUser(ctx context.Context, user *usermodel.User) (*uuid.UUID, error)
}

func (u *registerUseCase) Execute(ctx context.Context, user *usermodel.UserCreation) (*uuid.UUID, error) {
	exists, err := u.registerRepository.FindUserByEmail(ctx, user.Email)
	if err != nil && !errors.Is(err, usermodel.UserNotFound) {
		return nil, err
	}

	if exists != nil {
		return nil, usermodel.UserAlreadyExists
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	salt, err := u.hasher.GenSalt(16)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := u.hasher.Hash(user.Password, salt)
	if err != nil {
		return nil, err
	}

	data := &usermodel.User{
		ID:        id,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      usermodel.RoleUser,
		Password:  hashedPassword,
		Salt:      salt,
	}

	_, err = u.registerRepository.CreateNewUser(ctx, data)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
