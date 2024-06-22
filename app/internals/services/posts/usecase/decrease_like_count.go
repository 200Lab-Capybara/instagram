package postusecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	postsmodel "instagram/app/internals/services/posts/model"
	"instagram/common"
)

type decreaseLikeCountUseCase struct {
	postRepository DecreaseLikeCountRepository
}

func NewDecreaseLikeCountUseCase(postRepository DecreaseLikeCountRepository) DecreaseLikeCountUseCase {
	return &decreaseLikeCountUseCase{
		postRepository: postRepository,
	}
}

type DecreaseLikeCountRepository interface {
	GetByID(ctx context.Context, postID uuid.UUID) (*postsmodel.Post, error)
	DecreaseLikeCount(ctx context.Context, postID uuid.UUID) (bool, error)
}

type DecreaseLikeCountUseCase interface {
	Execute(ctx context.Context, postID uuid.UUID) (bool, error)
}

func (i *decreaseLikeCountUseCase) Execute(ctx context.Context, postID uuid.UUID) (bool, error) {
	post, err := i.postRepository.GetByID(ctx, postID)
	if err != nil {
		if errors.Is(err, postsmodel.ErrorPostNotFound) {
			return false, postsmodel.ErrorPostNotFound
		}
		return false, common.ErrInvalidRequest(err)
	}

	if post == nil {
		return false, postsmodel.ErrorPostNotFound
	}

	if post.LikeCount == 0 {
		return false, common.ErrInvalidRequest(errors.New("like count is already 0"))
	}

	ok, err := i.postRepository.DecreaseLikeCount(ctx, postID)
	if err != nil {
		return false, common.ErrInvalidRequest(err)
	}

	return ok, nil
}
