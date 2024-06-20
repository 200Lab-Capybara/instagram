package usermysql

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	usermodel "instagram/app/internals/services/user/model"
	"instagram/common"
)

func (s *mySQLStorage) FindUserById(ctx context.Context, id uuid.UUID) (*usermodel.User, error) {
	var data usermodel.User

	db := s.db.GetConnection().Table(data.TableName())

	if err := db.Where("id = ?", id).First(&data).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, usermodel.UserNotFound
		} else {
			return nil, common.ErrDB(err)
		}
	}

	return &data, nil
}
