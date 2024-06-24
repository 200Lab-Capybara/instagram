package mysqlreactcomment

import (
	"context"
	"github.com/google/uuid"
	modelreactioncomment "instagram/app/internals/services/reaction_comment/model"
	"instagram/common"
	"time"
)

func (m *mySQLStorage) CreateNewReactionComment(ctx context.Context, commentId uuid.UUID, userId uuid.UUID) (bool, error) {
	db := m.db.GetConnection()
	newRow := modelreactioncomment.ReactionComment{
		UserId:    userId,
		CommentId: commentId,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: nil,
	}
	if err := db.Table(modelreactioncomment.ReactionComment{}.TableName()).
		Create(newRow).Error; err != nil {
		return false, common.ErrDB(err)
	}
	return true, nil
}
