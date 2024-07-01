package storymysql

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"instagram/app/internals/services/stories/model"
	"instagram/common"
)

func (storage *mysqlStorage) GetById(ctx context.Context, storyId uuid.UUID) (*storiesmodel.Story, error) {
	story := storiesmodel.Story{}
	err := storage.db.GetConnection().WithContext(ctx).Where("id = ?", storyId).First(&story).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, storiesmodel.ErrStoryNotFound
		}

		return nil, common.ErrDB(err)
	}
	return &story, nil
}
