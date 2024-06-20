package hashtagsql

import (
	"context"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	hashtagmodel "instagram/app/internals/services/hashtag/model"
	"strings"
)

func FormatHashTag(ctx context.Context, hashtag []string) ([]hashtagmodel.Hashtag, error) {
	// Remove empty hashtags
	var nonNullHashtags []hashtagmodel.Hashtag
	for _, tag := range hashtag {
		if tag != "" {
			lowercaseTag := strings.ToLower(tag)
			sluggedTag := strings.ReplaceAll(slug.Make(lowercaseTag), " ", "-")
			validHashTag := hashtagmodel.Hashtag{ID: uuid.New(), Hashtag: sluggedTag}
			nonNullHashtags = append(nonNullHashtags, validHashTag)
		}
	}
	return nonNullHashtags, nil
}
