package followmysql

import (
	"context"
	"github.com/google/uuid"
	followusermodel "instagram/app/internals/services/follow/model"
)

func (store *mysqlStorage) GetListFollowingWithIds(ctx context.Context, uid uuid.UUID, ids []uuid.UUID) ([]followusermodel.FollowUser, error) {
	var data []followusermodel.FollowUser
	db := store.db.GetConnection()
	if err := db.Where("user_id = ? AND following IN ?", uid, ids).Table(followusermodel.FollowUser{}.TableName()).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}
