package userusecase

import (
	"context"
	usermodel "instagram/app/internals/services/user/model"
)

type updateFollowUseCase struct {
	updateFollowRepo UpdateFollowRepository
}

func NewUpdateFollowUseCase(updateFollowRepo UpdateFollowRepository) UpdateFollowUseCase {
	return &updateFollowUseCase{
		updateFollowRepo: updateFollowRepo,
	}
}

type UpdateFollowUseCase interface {
	Execute(ctx context.Context, dto *usermodel.FollowUserPayload) (bool, error)
}

type UpdateFollowRepository interface {
	UpdateFollow(ctx context.Context, dto *usermodel.FollowUserPayload) (bool, error)
}

func (u *updateFollowUseCase) Execute(ctx context.Context, dto *usermodel.FollowUserPayload) (bool, error) {
	_, err := u.updateFollowRepo.UpdateFollow(ctx, dto)
	if err != nil {
		return false, err
	}
	return true, nil
}
