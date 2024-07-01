package storyusecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"instagram/app/internals/services/reaction_story/model"
	storiesmodel "instagram/app/internals/services/stories/model"
	"instagram/common"
)

type increaseLikeCountUseCase struct {
	increaselikecount IncreaseLikeCountRepo
}

func NewIncreaseLikeCountUseCase(increaselikecount IncreaseLikeCountRepo) IncreaseLikeCountUseCase {
	return &increaseLikeCountUseCase{
		increaselikecount: increaselikecount,
	}
}

func (uc *increaseLikeCountUseCase) Execute(ctx context.Context, storyId uuid.UUID) (bool, error) {
	_, err := uc.increaselikecount.GetById(ctx, storyId)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return false, model.ErrStoryNotFound
		}
		return false, common.ErrDB(err)
	}
	ok, err := uc.increaselikecount.IncreaseLikeCount(ctx, storyId)
	if err != nil {
		return false, common.ErrInvalidRequest(err)
	}
	return ok, nil
}

type IncreaseLikeCountUseCase interface {
	Execute(ctx context.Context, storyId uuid.UUID) (bool, error)
}
type IncreaseLikeCountRepo interface {
	GetById(ctx context.Context, storyId uuid.UUID) (*storiesmodel.Story, error)
	IncreaseLikeCount(ctx context.Context, storyId uuid.UUID) (bool, error)
}
