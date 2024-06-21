package hashtagusercase

import (
	"context"
	"github.com/google/uuid"
	hashtagmodel "instagram/app/internals/services/hashtag/model"
	"time"
)

type addingHashTagUseCase struct {
	addingHashTagRepository AddingHashTagRepository
}

func NewCreateHashTagUseCase(addingHashTagRepository AddingHashTagRepository) AddingHashTagUseCase {
	return &addingHashTagUseCase{
		addingHashTagRepository: addingHashTagRepository,
	}
}

type AddingHashTagUseCase interface {
	Execute(ctx context.Context, postId uuid.UUID, hashtag []string) (bool, error)
}

type AddingHashTagRepository interface {
	MapHashTag(ctx context.Context, postID uuid.UUID, hashtag hashtagmodel.Hashtag) (bool, error)
}

func (u *addingHashTagUseCase) Execute(ctx context.Context, postId uuid.UUID, hashtags []string) (bool, error) {

	hashtagFormat := hashtagmodel.Hashtag{}
	validNullHashtags, err := hashtagFormat.HashTagFormat(ctx, hashtags)

	if err != nil {
		return false, err
	}

	for _, hashtag := range validNullHashtags {
		_, err := u.addingHashTagRepository.MapHashTag(ctx, postId, hashtagmodel.Hashtag{ID: hashtag.ID, Hashtag: hashtag.Hashtag, CreatedAt: time.Now().UTC()})
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
