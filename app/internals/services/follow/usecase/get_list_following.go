package followusecase

import (
	"context"
	"github.com/google/uuid"
	followusermodel "instagram/app/internals/services/follow/model"
	"instagram/common"
)

type getListFollowingUseCase struct {
	followRepo GetListFollowingRepository
	userRepo   GetListUserByIdsRepository
}

func NewGetListFollowingUseCase(followRepo GetListFollowingRepository, userRepo GetListUserByIdsRepository) GetListFollowingUseCase {
	return &getListFollowingUseCase{followRepo: followRepo, userRepo: userRepo}
}

type GetListFollowingUseCase interface {
	Execute(ctx context.Context, uid uuid.UUID, paging *common.Paging) ([]common.SimpleUser, error)
}

type GetListFollowingRepository interface {
	GetListFollowing(ctx context.Context, userId *uuid.UUID, paging *common.Paging) ([]followusermodel.FollowUser, error)
}

type GetListUserByIdsRepository interface {
	GetUserByIds(ctx context.Context, ids []uuid.UUID) ([]common.SimpleUser, error)
}

func (g *getListFollowingUseCase) Execute(ctx context.Context, uid uuid.UUID, paging *common.Paging) ([]common.SimpleUser, error) {
	paging.Process()
	userId := uid

	data, err := g.followRepo.GetListFollowing(ctx, &userId, paging)
	if err != nil {
		return nil, err
	}

	followingIds := make([]uuid.UUID, len(data))
	for i, follow := range data {
		followingIds[i] = follow.Following
	}

	users, err := g.userRepo.GetUserByIds(ctx, followingIds)
	if err != nil {
		return nil, err
	}

	mapUser := make(map[uuid.UUID]common.SimpleUser)
	for _, user := range users {
		mapUser[user.UserId] = user
	}

	result := make([]common.SimpleUser, len(data))
	for i, id := range followingIds {
		result[i] = mapUser[id]
	}

	return result, nil
}
