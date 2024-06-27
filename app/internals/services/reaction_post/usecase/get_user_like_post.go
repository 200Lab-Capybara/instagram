package reactionpostusecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	reactionpostmodel "instagram/app/internals/services/reaction_post/model"
	usermodel "instagram/app/internals/services/user/model"
	"instagram/common"
)

type getUserLikePostUseCase struct {
	getUserLikePostRepo GetUserLikePostRepo
	postRepository      GetPostRepository
	getUserInfoRepo     GetUserInfoRepo
	getListFollwingRepo GetListFollwingRepo
}

func GetUserLikePostUC(getUserLikePostRepo GetUserLikePostRepo, postRepository GetPostRepository, getUserInfoRepo GetUserInfoRepo, getListFollwingRepo GetListFollwingRepo) GetUserLikePostUseCase {
	return &getUserLikePostUseCase{
		getUserLikePostRepo: getUserLikePostRepo,
		postRepository:      postRepository,
		getUserInfoRepo:     getUserInfoRepo,
		getListFollwingRepo: getListFollwingRepo,
	}
}

func (uc *getUserLikePostUseCase) Execute(ctx context.Context, userId uuid.UUID, postId uuid.UUID) ([]reactionpostmodel.UserReactionPost, error) {
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

	listInfoUserLikePost := make([]common.SimpleUser, len(listUserId))

	for i, userId := range listUserId {

		listInfoUserLikePost[i] = common.SimpleUser{
			UserId:    userMap[userId].ID,
			FirstName: userMap[userId].FirstName,
			LastName:  userMap[userId].LastName,
		}

	}

	listUserIdFollwing, err := uc.getListFollwingRepo.GetListFollwingByUserId(ctx, userId)

	fmt.Println(listUserIdFollwing)

	listFollwingMap := make(map[uuid.UUID]bool)

	for _, userIdFoolwing := range listUserIdFollwing {
		listFollwingMap[userIdFoolwing] = true
	}

	listUserLikePost := make([]reactionpostmodel.UserReactionPost, len(listUserIdFollwing))

	for i, _ := range listUserIdFollwing {
		if listFollwingMap[listInfoUserLikePost[i].UserId] {
			listUserLikePost[i] = reactionpostmodel.UserReactionPost{
				SimpleUser: listInfoUserLikePost[i],
				Followed:   true,
			}
		}
	}

	return listUserLikePost, nil
}

type GetUserLikePostUseCase interface {
	Execute(ctx context.Context, userId uuid.UUID, postId uuid.UUID) ([]reactionpostmodel.UserReactionPost, error)
}

type GetUserLikePostRepo interface {
	GetUserIdLikePost(ctx context.Context, postId uuid.UUID) ([]uuid.UUID, error)
}

type GetUserInfoRepo interface {
	GetUserInfoById(ctx context.Context, listUserId []uuid.UUID) ([]usermodel.User, error)
}

type GetListFollwingRepo interface {
	GetListFollwingByUserId(ctx context.Context, userId uuid.UUID) ([]uuid.UUID, error)
}
