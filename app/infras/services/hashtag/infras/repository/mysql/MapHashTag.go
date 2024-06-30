package hashtagsql

import (
	"context"
	"github.com/google/uuid"
	hashtagmodel "instagram/app/internals/services/hashtag/model"
)

// MapHashTag maps a hashtag to a specific post in the database.
func (m *mySQLStorage) MapHashTag(ctx context.Context, postID uuid.UUID, hashtag hashtagmodel.Hashtag) (bool, error) {
	var data hashtagmodel.HashtagPost
	db := m.db.GetConnection()
	// Check if the hashtagPost is nil
	newHashTagPost := hashtagmodel.HashtagPost{Hashtag_ID: hashtag.ID, Post_ID: postID, CreatedAt: hashtag.CreatedAt}

	if err := db.Table(data.TableName()).Create(&newHashTagPost).Error; err != nil {
		return false, err
	}
	return true, nil
}
