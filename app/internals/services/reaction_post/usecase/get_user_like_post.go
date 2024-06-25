package reactionpostusecase

import (
	"context"
	"github.com/google/uuid"
	reactionpostmodel "instagram/app/internals/services/reaction_post/model"
	"instagram/common"
)

type getUserLikePostUseCase struct {
	getUserLikePostRepo GetUserLikePostRepo
	postRepository      GetPostRepository
}

func GetUserLikePostUC(getUserLikePostRepo GetUserLikePostRepo, postRepository GetPostRepository) GetUserLikePostUseCase {
	return &getUserLikePostUseCase{
		getUserLikePostRepo: getUserLikePostRepo,
		postRepository:      postRepository,
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

	listUser, err := uc.getUserLikePostRepo.ListingUserLikePost(ctx, postId)

	if err != nil {
		return false, err
	}

	return listUser, nil
}

type GetUserLikePostUseCase interface {
	Execute(ctx context.Context, post_id uuid.UUID) (any, error)
}

type GetUserLikePostRepo interface {
	ListingUserLikePost(ctx context.Context, postId uuid.UUID) ([]reactionpostmodel.ReactionPost, error)
}
