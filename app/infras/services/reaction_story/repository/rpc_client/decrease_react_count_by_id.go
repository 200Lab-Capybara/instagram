package rpc_client

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (repo *getStoryRepo) DecreaseReactCountById(ctx context.Context, sid uuid.UUID) (bool, error) {
	if err := repo.db.GetConnection().
		Table("stories").
		Where("id = ?", sid).
		UpdateColumn("react_count", gorm.Expr("react_count - ?", 1)).
		Error; err != nil {
		return false, err
	}

	return true, nil
}
