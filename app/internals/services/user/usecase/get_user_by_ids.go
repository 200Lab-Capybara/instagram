package userusecase

import (
	"context"
	"github.com/google/uuid"
	usermodel "instagram/app/internals/services/user/model"
)

type getUserByIdsUseCase struct {
	repo GetUserByIdsRepository
}

func NewGetUserByIdsUseCase(repo GetUserByIdsRepository) GetUserByIdsUseCase {
	return &getUserByIdsUseCase{repo: repo}
}

type GetUserByIdsRepository interface {
	FindUsersByIds(ctx context.Context, ids []uuid.UUID) ([]usermodel.User, error)
}

type GetUserByIdsUseCase interface {
	Execute(ctx context.Context, ids []uuid.UUID) ([]usermodel.User, error)
}

func (g getUserByIdsUseCase) Execute(ctx context.Context, ids []uuid.UUID) ([]usermodel.User, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	users, err := g.repo.FindUsersByIds(ctx, ids)
	if err != nil {
		return nil, err
	}

	return users, nil
}
