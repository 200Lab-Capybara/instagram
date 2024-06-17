package usermysql

import (
	"context"
	usermodel "github.com/nghiatrann0502/instagram-clone/app/internals/services/user/model"
	"gorm.io/gorm"
)

func (m *mySQLStorage) FindUserByEmail(ctx context.Context, email string) (*usermodel.User, error) {
	var data usermodel.User

	db := m.db.GetConnection().Table(data.TableName())
	if err := db.Where("email = ?", email).First(&data).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, usermodel.UserNotFound
		} else {
			return nil, err
		}
	}

	return &data, nil
}
