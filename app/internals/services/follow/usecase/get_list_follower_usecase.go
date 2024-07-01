package followusecase

import (
	"context"
	"github.com/google/uuid"
	followusermodel "instagram/app/internals/services/follow/model"
	"instagram/common"
)

type getListFollowerUseCase struct {
	userRepo   GetListFollowerUserByIdsRepository
	followRepo GetListFollowerRepository
}

func NewGetListFollowerUseCase(userRepo GetListFollowerUserByIdsRepository, followRepo GetListFollowerRepository) GetListFollowerUseCase {
	return &getListFollowerUseCase{
		userRepo:   userRepo,
		followRepo: followRepo,
	}
}

type GetListFollowerUseCase interface {
	Execute(ctx context.Context, userId uuid.UUID, paging *common.Paging) ([]followusermodel.Follower, error)
}

type GetListFollowerUserByIdsRepository interface {
	GetUserByIds(ctx context.Context, ids []uuid.UUID) ([]common.SimpleUser, error)
}

type GetListFollowerRepository interface {
	GetListFollowingWithIds(ctx context.Context, uid uuid.UUID, ids []uuid.UUID) ([]followusermodel.FollowUser, error)
	GetListFollower(ctx context.Context, userId uuid.UUID, paging *common.Paging) ([]followusermodel.FollowUser, error)
}

func (g *getListFollowerUseCase) Execute(ctx context.Context, userId uuid.UUID, paging *common.Paging) ([]followusermodel.Follower, error) {
	paging.Process()
	follower, err := g.followRepo.GetListFollower(ctx, userId, paging)
	if err != nil {
		return nil, err
	}

	followerIds := make([]uuid.UUID, len(follower))
	for i, follow := range follower {
		followerIds[i] = follow.UserID
	}

	users, err := g.userRepo.GetUserByIds(ctx, followerIds)
	if err != nil {
		return nil, err
	}

	mapUser := make(map[uuid.UUID]common.SimpleUser)
	for _, user := range users {
		mapUser[user.UserId] = user
	}

	listFollowing, err := g.followRepo.GetListFollowingWithIds(ctx, userId, followerIds)
	if err != nil {
		return nil, err
	}

	mapFollowing := make(map[uuid.UUID]followusermodel.FollowUser)
	for _, follow := range listFollowing {
		mapFollowing[follow.Following] = follow
	}

	result := make([]followusermodel.Follower, len(follower))
	for i, id := range followerIds {
		simpleUser := mapUser[id]
		result[i] = followusermodel.Follower{
			SimpleUser: simpleUser,
			Followed:   mapFollowing[id].UserID != uuid.Nil,
		}
	}

	return result, nil

}
