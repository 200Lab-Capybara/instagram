package hashtagusercase

import (
	"context"
	"github.com/google/uuid"
	hashtagmodel "github.com/nghiatrann0502/instagram-clone/app/internals/services/hashtag/model"
)
type createHashTagUseCase struct{
	createHashTagRepository CreateHashTagRepository
}

func NewCreateHashTagUseCase(createHashTagRepository CreateHashTagRepository) CreateHashTagUseCase{
	return &createHashTagUseCase{
		createHashTagRepository : createHashTagRepository,
	}
}

type CreateHashTagUseCase interface{
	Execute(ctx context.Context, postId uuid.UUID, hashtag []string) ([]*hashtagmodel.HashtagPost, error)
}

type CreateHashTagRepository interface{
	FormatHashTag(ctx context.Context, hashtag []string) ([]*hashtagmodel.Hashtag, error)
	// GetHashTag(ctx context.Context, hashtag []string) ([]*hashtagmodel.Hashtag, error)
	MapHashTag(ctx context.Context, hashtag *hashtagmodel.HashtagPost) (*hashtagmodel.HashtagPost,error)
}

func (u *createHashTagUseCase) Execute(ctx context.Context, postId uuid.UUID, hashtags []string) ([]*hashtagmodel.HashtagPost, error){
	
	nonNullHashtags, err := u.createHashTagRepository.FormatHashTag(ctx, hashtags)
	if err != nil {
		return nil, err
	}

	// Retrieve existing or create new hashtags
	// hashtagEntries, err := u.createHashTagRepository.GetHashTag(ctx, nonNullHashtags)
	// if err != nil {
	// 	return nil, err
	// }

	// Map each hashtag to the postId
	var mappedHashtags []*hashtagmodel.HashtagPost
	for _, hashtag := range nonNullHashtags {
		mappedHashtag, err := u.createHashTagRepository.MapHashTag(ctx, &hashtagmodel.HashtagPost{
			Hashtag_ID: hashtag.ID,
			Post_ID:    postId,
		})
		if err != nil {
			return nil, err
		}
		mappedHashtags = append(mappedHashtags, mappedHashtag)
	}

	return mappedHashtags, nil
}
