package hashtagmodel

import (
	"context"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"strings"
	"time"
)

type Hashtag struct {
	ID        uuid.UUID `json:"id"`
	Hashtag   string    `json:"hashtag"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Hashtag) TableName() string {
	return "hashtags"
}

func (Hashtag) HashTagFormat(ctx context.Context, hashtag []string) ([]string, error) {
	// Remove empty hashtags
	var nonNullHashtags []string
	for _, tag := range hashtag {
		if tag != "" {
			lowercaseTag := strings.ToLower(tag)
			sluggedTag := strings.ReplaceAll(slug.Make(lowercaseTag), " ", "-")
			//validHashTag := Hashtag{ID: uuid.New(), Hashtag: sluggedTag}
			nonNullHashtags = append(nonNullHashtags, sluggedTag)
		}
	}
	return nonNullHashtags, nil
}
func (Hashtag) HashTagConvert(ctx context.Context, hashtag []string) ([]Hashtag, error) {
	// Remove empty hashtags
	var hashtagSlice []Hashtag
	for _, tag := range hashtag {
		if tag != "" {
			validHashTag := Hashtag{ID: uuid.New(), Hashtag: tag}
			hashtagSlice = append(hashtagSlice, validHashTag)
		}
	}
	return hashtagSlice, nil
}
