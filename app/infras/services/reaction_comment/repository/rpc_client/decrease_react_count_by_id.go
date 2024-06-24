package rpc_client

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	modelreactioncomment "instagram/app/internals/services/reaction_comment/model"
)

func (repo *getCommentRepo) DecreaseReactCountById(ctx context.Context, commentId uuid.UUID) (bool, error) {
	if err := repo.db.GetConnection().
		Table(modelreactioncomment.Comment{}.TableName()).
		Where("id = ?", commentId).
		UpdateColumn("react_count", gorm.Expr("react_count - ?", 1)).Error; err != nil {
	}
	return true, nil
}
