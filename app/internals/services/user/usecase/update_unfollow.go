package userusecase

import (
	"context"
	usermodel "instagram/app/internals/services/user/model"
)

type updateUnfollowUseCase struct {
	unfollowRepo UpdateUnfollowRepository
}

func NewUpdateUnfollowUseCase(unfollowRepo UpdateUnfollowRepository) UpdateFollowUseCase {
	return &updateUnfollowUseCase{
		unfollowRepo: unfollowRepo,
	}
}

type UpdateUnfollowUseCase interface {
	Execute(ctx context.Context, dto *usermodel.FollowUserPayload) (bool, error)
}

type UpdateUnfollowRepository interface {
	UpdateUnfollow(ctx context.Context, dto *usermodel.FollowUserPayload) (bool, error)
}

func (u *updateUnfollowUseCase) Execute(ctx context.Context, dto *usermodel.FollowUserPayload) (bool, error) {
	_, err := u.unfollowRepo.UpdateUnfollow(ctx, dto)

	if err != nil {
		return false, err
	}

	return true, nil
}
