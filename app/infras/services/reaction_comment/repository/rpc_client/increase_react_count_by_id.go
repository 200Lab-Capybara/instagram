package rpc_clientreactioncomment

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	modelreactioncomment "instagram/app/internals/services/reaction_comment/model"
	"instagram/common"
)

func (repo *getCommentRepo) IncreaseReacCountById(ctx context.Context, commentId uuid.UUID) (bool, error) {
	if err := repo.db.GetConnection().
		Table(modelreactioncomment.Comment{}.TableName()).
		Where("id = ?", commentId).
		UpdateColumn("react_count", gorm.Expr("react_count + ?", 1)).
		Error; err != nil {
		return false, common.ErrDB(err)
	}
	return true, nil
}
