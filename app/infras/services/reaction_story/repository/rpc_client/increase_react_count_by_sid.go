package rpc_client

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"instagram/app/internals/services/reaction_story/model"
	"instagram/common"
)

func (repo *getStoryRepo) IncreaseReactCountById(ctx context.Context, sid uuid.UUID) (bool, error) {
	// Increment react_count by 1
	if err := repo.db.GetConnection().
		Table(model.Story{}.TableName()).
		Where("id = ?", sid).
		UpdateColumn("react_count", gorm.Expr("react_count + ?", 1)).
		Error; err != nil {
		return false, common.ErrDB(err)
	}

	return true, nil
}
