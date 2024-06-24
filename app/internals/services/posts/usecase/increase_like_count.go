package postusecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	postsmodel "instagram/app/internals/services/posts/model"
	"instagram/common"
)

type increaseLikeCountUseCase struct {
	postRepository IncreaseLikeCountRepository
}

func NewIncreaseLikeCountUseCase(postRepository IncreaseLikeCountRepository) IncreaseLikeCountUseCase {
	return &increaseLikeCountUseCase{
		postRepository: postRepository,
	}
}

type IncreaseLikeCountRepository interface {
	GetByID(ctx context.Context, postID uuid.UUID) (*postsmodel.Post, error)
	IncreaseLikeCount(ctx context.Context, postID uuid.UUID) (bool, error)
}

type IncreaseLikeCountUseCase interface {
	Execute(ctx context.Context, postID uuid.UUID) (bool, error)
}

func (i *increaseLikeCountUseCase) Execute(ctx context.Context, postID uuid.UUID) (bool, error) {
	post, err := i.postRepository.GetByID(ctx, postID)
	if err != nil {
		if errors.Is(err, postsmodel.ErrorPostNotFound) {
			return false, postsmodel.ErrorPostNotFound
		}
		return false, common.ErrInvalidRequest(err)
	}

	fmt.Println(post)

	if post == nil {
		return false, postsmodel.ErrorPostNotFound
	}

	ok, err := i.postRepository.IncreaseLikeCount(ctx, postID)
	if err != nil {
		return false, common.ErrInvalidRequest(err)
	}

	return ok, nil
}