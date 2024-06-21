package rpc_client

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (repo *getStoryRepo) IncreaseReactCountById(ctx context.Context, sid uuid.UUID) (bool, error) {
	// Increment react_count by 1
	if err := repo.db.GetConnection().
		Table("stories").
		Where("id = ?", sid).
		UpdateColumn("react_count", gorm.Expr("react_count + ?", 1)).
		Error; err != nil {
		return false, err
	}

	return true, nil
}
