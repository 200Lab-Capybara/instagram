package userusecase

import (
	"context"
	"github.com/google/uuid"
	usermodel "instagram/app/internals/services/user/model"
)

type getUserByIdUseCase struct {
	repo GetUserByIdRepository
}

func NewGetUserByIdUseCase(repo GetUserByIdRepository) GetUserByIdUseCase {
	return &getUserByIdUseCase{repo: repo}
}

type GetUserByIdUseCase interface {
	Execute(ctx context.Context, id uuid.UUID) (*usermodel.User, error)
}

type GetUserByIdRepository interface {
	FindUserById(ctx context.Context, id uuid.UUID) (*usermodel.User, error)
}

func (g *getUserByIdUseCase) Execute(ctx context.Context, id uuid.UUID) (*usermodel.User, error) {
	user, err := g.repo.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
