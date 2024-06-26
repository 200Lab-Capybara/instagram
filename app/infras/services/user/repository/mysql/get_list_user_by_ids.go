package usermysql

import (
	"context"
	"github.com/google/uuid"
	usermodel "instagram/app/internals/services/user/model"
	"instagram/common"
)

func (store *mySQLStorage) FindUsersByIds(ctx context.Context, ids []uuid.UUID) ([]usermodel.User, error) {
	db := store.db.GetConnection()
	var data []usermodel.User

	if err := db.Table(usermodel.User{}.TableName()).
		Where("id IN (?)", ids).
		Find(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return data, nil
}
