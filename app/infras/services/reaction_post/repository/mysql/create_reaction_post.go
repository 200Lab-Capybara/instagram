package reactionpostmysql

import (
	"context"
	"github.com/google/uuid"
	"instagram/app/internals/services/reaction_post/model"
	"time"
)

func (m *mySQLStorage) CreateNewReactionPost(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (bool, error) {
	db := m.db.GetConnection()

	newRow := model.ReactionPost{PostID: postId, UserID: userId, CreatedAt: time.Now().UTC()}

	if err := db.Table(model.ReactionPost{}.TableName()).
		Create(newRow).Error; err != nil {
		return false, err
	}
	return true, nil
}
