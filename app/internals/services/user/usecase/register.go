package userusecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	usermodel "instagram/app/internals/services/user/model"
	"instagram/common"
	"instagram/components/hasher"
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
	if err != nil && !errors.Is(err, usermodel.ErrUserNotFound) {
		return nil, err
	}

	if exists != nil {
		return nil, common.NewCustomError(usermodel.ErrUserAlreadyExists, usermodel.ErrUserAlreadyExists.Error(), "user_already_exists")
	}

	id, _ := uuid.NewV7()

	salt, _ := u.hasher.GenSalt(16)

	hashedPassword, err := u.hasher.Hash(user.Password, salt)

	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	data := &usermodel.User{
		ID:        id,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      common.RoleUser,
		Status:    common.UserActive,
		Password:  hashedPassword,
		Salt:      salt,
	}

	_, err = u.registerRepository.CreateNewUser(ctx, data)
	if err != nil {
		return nil, err
	}

	return &id, nil
}
