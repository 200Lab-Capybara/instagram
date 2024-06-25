package usermysql

import (
	"context"
	"gorm.io/gorm"
	usermodel "instagram/app/internals/services/user/model"
	"instagram/common"
)

func (store *mySQLStorage) UpdateUnfollow(ctx context.Context, dto *usermodel.FollowUserPayload) (bool, error) {
	db := store.db.GetConnection()

	if err := db.Table(dto.TableName()).Where("id = ?", dto.FollowingID).Update("follower", gorm.Expr("follower - ?", 1)).Error; err != nil {
		return false, common.ErrDB(err)
	}

	if err := db.Table(dto.TableName()).Where("id = ?", dto.UserID).Update("following", gorm.Expr("following - ?", 1)).Error; err != nil {
		return false, common.ErrDB(err)
	}

	return true, nil
}
