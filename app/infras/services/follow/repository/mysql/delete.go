package followmysql

import (
	"context"
	"github.com/google/uuid"
	followusermodel "instagram/app/internals/services/follow/model"
	"instagram/common"
)

func (store *mysqlStorage) Delete(ctx context.Context, follower uuid.UUID, following uuid.UUID) (bool, error) {
	db := store.db.GetConnection()
	data := followusermodel.FollowUser{}

	if err := db.Table(data.TableName()).
		Where("user_id = ? AND following = ?", follower, following).
		Delete(&data).Error; err != nil {
		return false, common.ErrDB(err)
	}

	return true, nil
}
