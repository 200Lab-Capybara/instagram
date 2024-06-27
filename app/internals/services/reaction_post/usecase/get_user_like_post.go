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

func (uc *getUserLikePostUseCase) Execute(ctx context.Context, postId uuid.UUID) ([]common.SimpleUser, error) {
	post, err := uc.postRepository.FindById(ctx, postId)
	if err != nil {
		return nil, err
	}
	if post.Status == "deleted" {
		return nil, common.ErrInvalidRequest(reactionpostmodel.ErrPostDoNotExist)
	}

	listUserId, err := uc.getUserLikePostRepo.GetUserIdLikePost(ctx, postId)
	if err != nil {
		return nil, err
	}

	if len(listUserId) == 0 {
		return nil, common.ErrInvalidRequest(reactionpostmodel.ErrRecordReactPostNotFound)
	}

	listInfo, err := uc.getUserInfoRepo.GetUserInfoById(ctx, listUserId)
	if err != nil {
		return nil, err
	}

	userMap := make(map[uuid.UUID]usermodel.User)
	for _, userInfo := range listInfo {
		userMap[userInfo.ID] = userInfo
	}

	listUserLikePost := make([]common.SimpleUser, len(listUserId))

	for i, userId := range listUserId {

		listUserLikePost[i] = common.SimpleUser{
			UserId:    userMap[userId].ID,
			FirstName: userMap[userId].FirstName,
			LastName:  userMap[userId].LastName,
		}

	}

	return listUserLikePost, nil
}

type GetUserLikePostUseCase interface {
	Execute(ctx context.Context, post_id uuid.UUID) ([]common.SimpleUser, error)
}

type GetUserLikePostRepo interface {
	GetUserIdLikePost(ctx context.Context, postId uuid.UUID) ([]uuid.UUID, error)
}

type GetUserInfoRepo interface {
	GetUserInfoById(ctx context.Context, listUserId []uuid.UUID) ([]usermodel.User, error)
}
