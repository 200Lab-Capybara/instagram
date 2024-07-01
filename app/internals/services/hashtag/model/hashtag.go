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

//func (Hashtag) HashTagFormat(ctx context.Context, hashtag []string) ([]string, error) {
//	// Remove empty hashtags
//	var nonNullHashtags []string
//	for _, tag := range hashtag {
//		if tag != "" {
//			lowercaseTag := strings.ToLower(tag)
//			sluggedTag := strings.ReplaceAll(slug.Make(lowercaseTag), " ", "-")
//			nonNullHashtags = append(nonNullHashtags, sluggedTag)
//		}
//	}
//
//	return nonNullHashtags, nil
//}

func (Hashtag) HashTagFormat(ctx context.Context, hashtags []string) ([]string, error) {
	// Preallocate slice with the same capacity as the input to avoid resizing
	nonNullHashtags := make([]string, 0, len(hashtags))
	index := 0 // Index for inserting non-empty hashtags into the slice

	for _, tag := range hashtags {
		if tag != "" {
			lowercaseTag := strings.ToLower(tag)
			sluggedTag := strings.ReplaceAll(slug.Make(lowercaseTag), " ", "-")
			nonNullHashtags = nonNullHashtags[:index+1]
			nonNullHashtags[index] = sluggedTag
			index++
		}
	}
	return nonNullHashtags[:index], nil
}

//func (Hashtag) HashTagConvert(ctx context.Context, hashtag []string) ([]Hashtag, error) {
//	// Remove empty hashtags
//	var hashtagSlice []Hashtag
//	for _, tag := range hashtag {
//		if tag != "" {
//			validHashTag := Hashtag{ID: uuid.New(), Hashtag: tag}
//			hashtagSlice = append(hashtagSlice, validHashTag)
//		}
//	}
//	return hashtagSlice, nil
//}

func (Hashtag) HashTagConvert(ctx context.Context, hashtags []string) ([]Hashtag, error) {
	// Preallocate slice with the maximum possible size to avoid reallocation
	hashtagSlice := make([]Hashtag, 0, len(hashtags))
	index := 0 // Index for inserting non-empty hashtags into the slice

	for _, tag := range hashtags {
		if tag != "" {
			newHashtag := Hashtag{
				ID:      uuid.New(), // Generate a new UUID for each hashtag
				Hashtag: tag,        // Use the non-empty string tag
			}
			hashtagSlice = hashtagSlice[:index+1] // Ensure the slice includes the new element
			hashtagSlice[index] = newHashtag      // Directly assign the new Hashtag struct
			index++                               // Increment the index for the next insertion
		}
	}

	return hashtagSlice[:index], nil // Return the slice up to the last filled index
}
