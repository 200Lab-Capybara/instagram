package followmysql

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	followusermodel "instagram/app/internals/services/follow/model"
	"instagram/common"
)

func (store *mysqlStorage) CheckIsExist(ctx context.Context, uId, fId uuid.UUID) (bool, error) {
	db := store.db.GetConnection()
	data := followusermodel.FollowUser{}

	if err := db.Table(data.TableName()).Where("user_id = ? AND following = ?", uId, fId).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, common.ErrDB(err)
	}

	return true, nil
}
