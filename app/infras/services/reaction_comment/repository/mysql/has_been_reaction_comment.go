package mysqlreactcomment

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	modelreactioncomment "instagram/app/internals/services/reaction_comment/model"
	"instagram/common"
)

func (m *mySQLStorage) HasBeenReactionComment(ctx context.Context, commentId uuid.UUID, userId uuid.UUID) (bool, error) {
	db := m.db.GetConnection().Debug()

	newRow := modelreactioncomment.ReactionComment{}

	err := db.Table(modelreactioncomment.ReactionComment{}.TableName()).
		Where("user_id = ? AND comment_id = ?", userId, commentId).
		First(&newRow).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, modelreactioncomment.ErrCommentNotFound
		}
		return false, common.ErrDB(err)
	}
	return true, nil
}
