package hashtagusercase

import (
	"context"
	"github.com/google/uuid"
	hashtagmodel "instagram/app/internals/services/hashtag/model"
)

type creatingHashTagPostUseCase struct {
	creatingHashTagPostRepository CreatingHashTagPostRepository
	creatingHashTagRepository     CreatingHashTagRepository
}

func NewCreateHashTagUseCase(creatingHashTagPostRepository CreatingHashTagPostRepository, creatingHashTagRepository CreatingHashTagRepository) CreatingHashTagPostUseCase {
	return &creatingHashTagPostUseCase{
		creatingHashTagPostRepository: creatingHashTagPostRepository,
		creatingHashTagRepository:     creatingHashTagRepository,
	}
}

type CreatingHashTagPostUseCase interface {
	Execute(ctx context.Context, postId uuid.UUID, hashtag []string) (bool, error)
}

type CreatingHashTagPostRepository interface {
	MapHashTag(ctx context.Context, postID uuid.UUID, hashtags []hashtagmodel.Hashtag) (bool, error)
}

type CreatingHashTagRepository interface {
	CreateHashTag(ctx context.Context, newHashTag []hashtagmodel.Hashtag) (bool, error)
	//GetHashTags(ctx context.Context, hashtagSlice []hashtagmodel.Hashtag) ([]*hashtagmodel.Hashtag, error)
	GetHashTags(ctx context.Context, hashtagSlice []string) ([]*hashtagmodel.Hashtag, error)
}

func (u *creatingHashTagPostUseCase) Execute(ctx context.Context, postId uuid.UUID, hashtags []string) (bool, error) {

	hashtagFormat := hashtagmodel.Hashtag{}
	validNullHashtagsString, err := hashtagFormat.HashTagFormat(ctx, hashtags)
	if err != nil {
		return false, err
	}

	existedHashtags, err := u.creatingHashTagRepository.GetHashTags(ctx, validNullHashtagsString)
	if err != nil {
		return false, err

	}

	validNullHashtag, err := hashtagFormat.HashTagConvert(ctx, validNullHashtagsString)
	if err != nil {
		return false, err
	}

	hashtagsStatusMap := make(map[string]bool)

	for _, hashtagString := range hashtags {
		hashtagsStatusMap[hashtagString] = false
	}

	if len(existedHashtags) != 0 {
		for _, hashtag := range existedHashtags {
			hashtagsStatusMap[hashtag.Hashtag] = true
		}
	}

	//fmt.Println("Hashtag Status Map:")
	//for hashtag, status := range hashtagsStatusMap {
	//	fmt.Printf("Hashtag: %s, Exists in DB: %t\n", hashtag, status)
	//}

	//for _, newHashTag := range validNullHashtag {
	//	if hashtagsStatusMap[newHashTag.Hashtag] != true {
	//		// creating hashtag
	//		_, creatingHashtagErr := u.creatingHashTagRepository.CreateHashTag(ctx, newHashTag)
	//		if creatingHashtagErr != nil {
	//			return false, creatingHashtagErr
	//		}
	//
	//		// mapping hashtag to post
	//		_, MappingErr := u.creatingHashTagPostRepository.MapHashTag(ctx, postId, hashtagmodel.Hashtag{ID: newHashTag.ID, Hashtag: newHashTag.Hashtag, CreatedAt: time.Now().UTC()})
	//		if MappingErr != nil {
	//			return false, MappingErr
	//		}
	//		hashtagsStatusMap[newHashTag.Hashtag] = true
	//	}
	//}

	newHashtag := make([]hashtagmodel.Hashtag, 0, len(hashtags))
	index := 0
	for _, newHashTagObj := range validNullHashtag {
		if hashtagsStatusMap[newHashTagObj.Hashtag] != true {
			newHashtag = newHashtag[:index+1]
			newHashtag[index] = newHashTagObj
			index++
		}
	}

	_, creatingHashtagErr := u.creatingHashTagRepository.CreateHashTag(ctx, newHashtag)
	if creatingHashtagErr != nil {
		return false, creatingHashtagErr
	}

	_, MappingErr := u.creatingHashTagPostRepository.MapHashTag(ctx, postId, newHashtag)
	if MappingErr != nil {
		return false, MappingErr
	}

	return true, nil
}
