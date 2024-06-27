package followmysql

import (
	"context"
	followusermodel "instagram/app/internals/services/follow/model"
	"instagram/common"
)

func (store *mysqlStorage) Create(ctx context.Context, dto *followusermodel.FollowUser) (bool, error) {
	db := store.db.GetConnection()

	if err := db.Table(dto.TableName()).Create(&dto).Error; err != nil {
		return false, common.ErrDB(err)
	}

	return true, nil
}
