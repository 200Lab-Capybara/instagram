package userinterface

import (
	"context"
	"github.com/google/uuid"
	usermodel "github.com/nghiatrann0502/instagram-clone/app/internals/services/user/model"
)

type UserRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*usermodel.User, error)
	CreateNewUser(ctx context.Context, user *usermodel.User) (*uuid.UUID, error)
}

type UserUseCase interface {
	Register(ctx context.Context, user *usermodel.UserCreation) (*uuid.UUID, error)
}
