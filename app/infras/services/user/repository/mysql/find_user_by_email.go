package usermysql

import (
	"context"
	"gorm.io/gorm"
	usermodel "instagram/app/internals/services/user/model"
	"instagram/common"
)

func (m *mySQLStorage) FindUserByEmail(ctx context.Context, email string) (*usermodel.User, error) {
	var data usermodel.User

	db := m.db.GetConnection().Table(data.TableName())
	if err := db.Where("email = ?", email).First(&data).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, usermodel.ErrUserNotFound
		} else {
			return nil, common.ErrDB(err)
		}
	}

	return &data, nil
}
