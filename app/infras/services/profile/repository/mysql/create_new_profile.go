package mysql

import (
	"context"
	"github.com/google/uuid"
	"instagram/app/internals/services/profile/model"
	"instagram/common"
)

func (m *mySQLStorage) CreateNewProfile(ctx context.Context, profile *model.Profile) (*uuid.UUID, error) {

	db := m.db.GetConnection()
	if err := db.Create(profile).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &profile.ID, nil
}
