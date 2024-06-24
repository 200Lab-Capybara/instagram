package mysqlreactcomment

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	modelreactioncomment "instagram/app/internals/services/reaction_comment/model"
	"instagram/common"
	"time"
)

func (m *mySQLStorage) HashBeenReactionComment(ctx context.Context, commentId uuid.UUID, userId uuid.UUID) (bool, error) {
	db := m.db.GetConnection()
	newRow := &modelreactioncomment.ReactionComment{
		UserId:    userId,
		CommentId: commentId,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: nil,
	}
	err := db.Table(modelreactioncomment.ReactionComment{}.TableName()).
		Where("user_id AND comment_id = ?", userId, commentId).
		First(newRow).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, modelreactioncomment.ErrCommentNotFound
		}
		return false, common.ErrDB(err)
	}
	return true, nil
}
