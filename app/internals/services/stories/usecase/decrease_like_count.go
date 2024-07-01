package storyusecase

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	storiesmodel "instagram/app/internals/services/stories/model"
	"instagram/common"
)

type decreaseLikeCountUseCase struct {
	decrease DecreaseLikeCountRepo
}

func NewDecreaseLikeCountUseCase(decreaselikecount DecreaseLikeCountRepo) DecreaseLikeCountUseCase {
	return &decreaseLikeCountUseCase{
		decrease: decreaselikecount,
	}
}
func (uc *decreaseLikeCountUseCase) Execute(ctx context.Context, storyId uuid.UUID) (bool, error) {
	_, err := uc.decrease.GetById(ctx, storyId)
	if err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return false, storiesmodel.ErrStoryNotFound
		}
		return false, common.ErrDB(err)
	}
	ok, err := uc.decrease.DecreaseLikeCount(ctx, storyId)
	if err != nil {
		return false, common.ErrInvalidRequest(err)
	}

	return ok, nil
}

type DecreaseLikeCountUseCase interface {
	Execute(ctx context.Context, storyId uuid.UUID) (bool, error)
}
type DecreaseLikeCountRepo interface {
	GetById(ctx context.Context, storyId uuid.UUID) (*storiesmodel.Story, error)
	DecreaseLikeCount(ctx context.Context, storyId uuid.UUID) (bool, error)
}
