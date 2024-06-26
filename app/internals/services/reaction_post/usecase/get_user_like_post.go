package reactionpostusecase

import (
	"context"
	"github.com/google/uuid"
	reactionpostmodel "instagram/app/internals/services/reaction_post/model"
	usermodel "instagram/app/internals/services/user/model"
	"instagram/common"
)

type getUserLikePostUseCase struct {
	getUserLikePostRepo GetUserLikePostRepo
	postRepository      GetPostRepository
	getUserInfoRepo     GetUserInfoRepo
}

func GetUserLikePostUC(getUserLikePostRepo GetUserLikePostRepo, postRepository GetPostRepository, getUserInfoRepo GetUserInfoRepo) GetUserLikePostUseCase {
	return &getUserLikePostUseCase{
		getUserLikePostRepo: getUserLikePostRepo,
		postRepository:      postRepository,
		getUserInfoRepo:     getUserInfoRepo,
	}
}

func (uc *getUserLikePostUseCase) Execute(ctx context.Context, postId uuid.UUID) (any, error) {
	post, err := uc.postRepository.FindById(ctx, postId)
	if err != nil {
		return false, err
	}
	if post.Status == "deleted" {
		return false, common.ErrInvalidRequest(reactionpostmodel.ErrPostDoNotExist)
	}

	listUser, err := uc.getUserLikePostRepo.GetUserIdLikePost(ctx, postId)
	if err != nil {
		return false, err
	}

	if len(listUser) == 0 {
		return false, common.ErrInvalidRequest(reactionpostmodel.ErrRecordReactPostNotFound)
	}

	listInfo, err := uc.getUserInfoRepo.GetUserInfoById(ctx, listUser)
	if err != nil {
		return false, err
	}

	userMap := make(map[uuid.UUID]usermodel.User)
	for _, userInfo := range listInfo {
		userMap[userInfo.ID] = userInfo
	}

	return userMap, nil
}

type GetUserLikePostUseCase interface {
	Execute(ctx context.Context, post_id uuid.UUID) (any, error)
}

type GetUserLikePostRepo interface {
	GetUserIdLikePost(ctx context.Context, postId uuid.UUID) ([]uuid.UUID, error)
}

type GetUserInfoRepo interface {
	GetUserInfoById(ctx context.Context, listUserId []uuid.UUID) ([]usermodel.User, error)
}
