package usermysql

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	usermodel "instagram/app/internals/services/user/model"
	"instagram/common"
)

func (store *mySQLStorage) FindUserById(ctx context.Context, id uuid.UUID) (*usermodel.User, error) {
	var data usermodel.User

	db := store.db.GetConnection().Table(data.TableName())

	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, usermodel.ErrUserNotFound
		} else {
			return nil, common.ErrDB(err)
		}
	}

	return &data, nil
}
