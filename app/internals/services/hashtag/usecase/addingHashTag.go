package hashtagusercase

import (
	"context"
	"github.com/google/uuid"
	hashtagmodel "instagram/app/internals/services/hashtag/model"
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
	FormatHashTag(ctx context.Context, hashtag []string) ([]*hashtagmodel.Hashtag, error)
	MapHashTag(ctx context.Context, hashtag *hashtagmodel.HashtagPost) (bool, error)
}

func (u *addingHashTagUseCase) Execute(ctx context.Context, postId uuid.UUID, hashtags []string) (bool, error) {

	validNullHashtags, err := u.addingHashTagRepository.FormatHashTag(ctx, hashtags)
	if err != nil {
		return false, err
	}

	for _, hashtag := range validNullHashtags {
		_, err := u.addingHashTagRepository.MapHashTag(ctx, &hashtagmodel.HashtagPost{Hashtag_ID: hashtag.ID, Post_ID: postId, CreatedAt: hashtag.CreatedAt})
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
