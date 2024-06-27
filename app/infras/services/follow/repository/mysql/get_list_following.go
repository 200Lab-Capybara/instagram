package followmysql

import (
	"context"
	"github.com/google/uuid"
	followusermodel "instagram/app/internals/services/follow/model"
	"instagram/common"
)

func (store *mysqlStorage) GetInternalListFollowing(ctx context.Context, userId *uuid.UUID) ([]uuid.UUID, error) {
	var data []followusermodel.FollowUser
	db := store.db.GetConnection()

	db = db.Table(followusermodel.FollowUser{}.TableName()).Where("user_id = ?", &userId)

	if err := db.Select("*").
		Order("created_at desc").
		Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	flIds := make([]uuid.UUID, len(data))
	for i, v := range data {
		flIds[i] = v.Following
	}

	return flIds, nil
}
