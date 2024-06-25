package reactionpostmysql

import (
	"context"
	"github.com/google/uuid"
	reactionpostmodel "instagram/app/internals/services/reaction_post/model"
	"instagram/common"
)

func (store *mySQLStorage) GetUserIdLikePost(ctx context.Context, postId uuid.UUID) ([]uuid.UUID, error) {
	db := store.db.GetConnection()
	var userIds []uuid.UUID
	if err := db.Table(reactionpostmodel.ReactionPost{}.TableName()).
		Where("post_id = ?", postId).Pluck("user_id", &userIds).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return userIds, nil
}
