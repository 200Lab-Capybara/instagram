package userusecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	usermodel "github.com/nghiatrann0502/instagram-clone/app/internals/services/user/model"
)

func (u *useCase) Register(ctx context.Context, user *usermodel.UserCreation) (*uuid.UUID, error) {
	exists, err := u.userRepository.FindUserByEmail(ctx, user.Email)
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
		Role:      usermodel.RoleAdmin,
		Password:  hashedPassword,
		Salt:      salt,
	}

	_, err = u.userRepository.CreateNewUser(ctx, data)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
