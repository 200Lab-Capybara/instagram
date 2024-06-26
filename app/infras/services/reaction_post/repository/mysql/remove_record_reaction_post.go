package reactionpostmysql

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	reactionpostmodel "instagram/app/internals/services/reaction_post/model"
	"instagram/common"
	"log"
)

func (m *mySQLStorage) RemoveRecordReactionPost(ctx context.Context, postId uuid.UUID, userId uuid.UUID) (bool, error) {
	db := m.db.GetConnection()

	if err := db.Table(reactionpostmodel.ReactionPost{}.TableName()).
		Where("post_id = ? AND user_id = ?", postId, userId).
		Delete(&reactionpostmodel.ReactionPost{}).Error; err != nil {
		log.Println(err, "err")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, reactionpostmodel.ErrRecordReactPostNotFound
		}
		return false, common.ErrDB(err)
	}

	return true, nil
}
