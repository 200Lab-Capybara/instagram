package mysqlreactcomment

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	modelreactioncomment "instagram/app/internals/services/reaction_comment/model"
	"instagram/common"
)

func (m *mySQLStorage) RemoveReactionComment(ctx context.Context, commentId uuid.UUID, userId uuid.UUID) (bool, error) {
	db := m.db.GetConnection()
	if err := db.Table(modelreactioncomment.ReactionComment{}.TableName()).
		Where("comment_id = ? AND user_id = ?", commentId, userId).
		Delete(&modelreactioncomment.ReactionComment{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, modelreactioncomment.ErrCommentNotFound
		}
		return false, common.ErrDB(err)
	}
	return true, nil
}
