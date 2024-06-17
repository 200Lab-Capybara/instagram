package usermysql

import (
	"context"
	"github.com/google/uuid"
	usermodel "github.com/nghiatrann0502/instagram-clone/app/internals/services/user/model"
)

func (m *mySQLStorage) CreateNewUser(ctx context.Context, user *usermodel.User) (*uuid.UUID, error) {

	db := m.db.GetConnection()
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	return &user.ID, nil
}
