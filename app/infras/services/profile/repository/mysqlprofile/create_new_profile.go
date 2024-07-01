package mysqlprofile

import (
	"context"
	"github.com/google/uuid"
	"instagram/app/internals/services/profile/model"
	"instagram/common"
)

func (m *mySQLStorage) InsertProfile(ctx context.Context, profile *model.Profile) (*uuid.UUID, error) {

	db := m.db.GetConnection()
	if err := db.Table(profile.TableName()).Create(profile).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &profile.ID, nil
}
