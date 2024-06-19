package hashtagsql

import (
	"context"
	"github.com/gosimple/slug"
	"strings"
)

func FormatHashTag(ctx context.Context, hashtag []string) ([]string, error) {
	// Remove empty hashtags
	var nonNullHashtags []string
	for _, tag := range hashtag {
		if tag != "" {
			lowercaseTag := strings.ToLower(tag)
			sluggedTag := strings.ReplaceAll(slug.Make(lowercaseTag), " ", "-")
			nonNullHashtags = append(nonNullHashtags, sluggedTag)
		}
	}
	return nonNullHashtags, nil
}
