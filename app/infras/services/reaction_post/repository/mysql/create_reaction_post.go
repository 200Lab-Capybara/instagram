package reactionpostmysql

import (
	"context"
	"github.com/google/uuid"
	reactionpostmodel "instagram/app/internals/services/reaction_post/model"
	"time"
)

func (m *mySQLStorage) CreateNewReactionPost(ctx context.Context, userId uuid.UUID, postId uuid.UUID) (bool, error) {
	db := m.db.GetConnection()

	newRow := reactionpostmodel.ReactionPost{PostID: postId, UserID: userId, CreatedAt: time.Now().UTC()}

	if err := db.Table(reactionpostmodel.ReactionPost{}.TableName()).
		Create(newRow).Error; err != nil {
		return false, err
	}
	return true, nil
}
