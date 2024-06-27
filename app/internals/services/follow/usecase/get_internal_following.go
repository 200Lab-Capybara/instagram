package followusecase

import (
	"context"
	"github.com/google/uuid"
)

type getInternalFollowingUseCase struct {
	followRepo GetInternalFollowingRepository
}

func NewGetInternalFollowingUseCase(followRepo GetInternalFollowingRepository) GetInternalFollowingUseCase {
	return &getInternalFollowingUseCase{followRepo: followRepo}
}

type GetInternalFollowingUseCase interface {
	Execute(ctx context.Context, uid *uuid.UUID) ([]uuid.UUID, error)
}

type GetInternalFollowingRepository interface {
	GetInternalListFollowing(ctx context.Context, userId *uuid.UUID) ([]uuid.UUID, error)
}

func (g getInternalFollowingUseCase) Execute(ctx context.Context, uid *uuid.UUID) ([]uuid.UUID, error) {
	data, err := g.followRepo.GetInternalListFollowing(ctx, uid)
	if err != nil {
		return nil, err
	}

	return data, err
}
