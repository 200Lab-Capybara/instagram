package reactionpostmysql

import (
	"context"
	"github.com/google/uuid"
	reactionpostmodel "instagram/app/internals/services/reaction_post/model"
	"instagram/common"
)

func (store *mySQLStorage) ListingUserLikePost(ctx context.Context, postId uuid.UUID) ([]reactionpostmodel.ReactionPost, error) {
	db := store.db.GetConnection()
	var reactions []reactionpostmodel.ReactionPost
	if err := db.Table(reactionpostmodel.ReactionPost{}.TableName()).
		Where("post_id = ?", postId).Find(&reactions).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return reactions, nil
}
