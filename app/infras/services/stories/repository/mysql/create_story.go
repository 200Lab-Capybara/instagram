package storymysql

import (
	"context"
	"github.com/google/uuid"
	"instagram/app/internals/services/stories/model"
	"instagram/common"
)

func (s *mysqlStorage) CreateStory(ctx context.Context, story *model.Story) (*uuid.UUID, error) {
	db := s.db.GetConnection()
	if err := db.Table(story.TableName()).Create(&story).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &story.Id, nil
}
