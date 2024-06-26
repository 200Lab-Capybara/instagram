package followmysql

import (
	"context"
	"github.com/google/uuid"
	followusermodel "instagram/app/internals/services/follow/model"
	"instagram/common"
)

func (store *mysqlStorage) GetListFollower(ctx context.Context, userId uuid.UUID, paging *common.Paging) ([]followusermodel.FollowUser, error) {
	var data []followusermodel.FollowUser
	db := store.db.GetConnection()

	db = db.Table(followusermodel.FollowUser{}.TableName()).Where("following = ?", userId)

	if err := db.Select("user_id").Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if v := paging.Cursor; v != "" {
		db = db.Where("created_at < ?", v)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Select("*").
		Order("created_at desc").
		Limit(paging.Limit).
		Find(&data).Error; err != nil {

		return nil, common.ErrDB(err)
	}

	if len(data) > 0 {
		paging.NextCursor = data[len(data)-1].CreatedAt.String()
	}

	return data, nil
}
