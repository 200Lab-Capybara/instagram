package reactionpostusecase

import (
	"context"
	"github.com/google/uuid"
	"instagram/app/internals/services/reaction_post/model"
)

type reactionPostUseCase struct {
	likePostRepository LikePostRepository
	postRepository     GetPostRepository
}

func NewLikePostUseCase(likePostRepository LikePostRepository, postRepository GetPostRepository) LikePostUseCase {
	return &reactionPostUseCase{
		likePostRepository: likePostRepository,
		postRepository:     postRepository,
	}
}

func (uc *reactionPostUseCase) Execute(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (bool, error) {
	post, err := uc.postRepository.FindById(ctx, postId)
	if err != nil {
		return false, err
	}

	if post.Status == "deleted" {

	}

	_, err = uc.likePostRepository.CreateNewReactionPost(ctx, userId, postId)
	if err != nil {
		return false, err
	}

	return true, nil
}

type LikePostUseCase interface {
	Execute(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (bool, error)
}

type LikePostRepository interface {
	CreateNewReactionPost(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (bool, error)
	RemoveRecordReactionPost(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (bool, error)
}

type GetPostRepository interface {
	FindById(ctx context.Context, postId uuid.UUID) (*model.Post, error)
}
