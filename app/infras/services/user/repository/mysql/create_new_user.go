package usermysql

import (
	"context"
	"github.com/google/uuid"
	usermodel "instagram/app/internals/services/user/model"
	"instagram/common"
)

func (store *mySQLStorage) CreateNewUser(ctx context.Context, user *usermodel.User) (*uuid.UUID, error) {

	db := store.db.GetConnection()
	if err := db.Table(user.TableName()).Create(user).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &user.ID, nil
}
