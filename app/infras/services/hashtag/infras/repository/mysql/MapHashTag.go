package hashtagsql

import (
	"context"
	"github.com/google/uuid"
	hashtagmodel "instagram/app/internals/services/hashtag/model"
)

// MapHashTag maps a hashtag to a specific post in the database.
//func (m *mySQLStorage) MapHashTag(ctx context.Context, postID uuid.UUID, hashtag []hashtagmodel.Hashtag) (bool, error) {
//	var data hashtagmodel.HashtagPost
//	db := m.db.GetConnection()
//	// Check if the hashtagPost is nil
//
//	newHashTagPost := hashtagmodel.HashtagPost{Hashtag_ID: hashtag.ID, Post_ID: postID, CreatedAt: hashtag.CreatedAt}
//
//	if err := db.Table(data.TableName()).Create(&newHashTagPost).Error; err != nil {
//		return false, err
//	}
//	return true, nil
//}

func (m *mySQLStorage) MapHashTag(ctx context.Context, postID uuid.UUID, hashtags []hashtagmodel.Hashtag) (bool, error) {
	// Preallocate slice for all mappings to be inserted
	var data hashtagmodel.HashtagPost
	db := m.db.GetConnection()
	hashtagPosts := make([]hashtagmodel.HashtagPost, len(hashtags))

	// Populate the slice using direct indexing
	for i, hashtag := range hashtags {
		hashtagPosts[i] = hashtagmodel.HashtagPost{
			Hashtag_ID: hashtag.ID,
			Post_ID:    postID,
			CreatedAt:  hashtag.CreatedAt,
		}
	}

	// Perform the batch insert using GORM
	if err := db.Table(data.TableName()).Create(&hashtagPosts).Error; err != nil {
		return false, err // Return false and the error if the batch insert fails
	}

	return true, nil // Return true if the batch insert is successful
}
