package reactionpostusecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/nghiatrann0502/instagram-clone/app/internals/services/reaction_post/model"
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

func (r *reactionPostUseCase) Execute(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (bool, error) {
	post, err := r.postRepository.FindById(ctx, postId)
	if err != nil {
		return false, err
	}

	if post.Status == "deleted" {

	}

	_, err = r.likePostRepository.CreateNewReactionPost(ctx, userId, postId)
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
}

type GetPostRepository interface {
	FindById(ctx context.Context, postId uuid.UUID) (*model.Post, error)
}
