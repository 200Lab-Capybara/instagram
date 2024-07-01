package reactstorymysql

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"instagram/app/internals/services/reaction_story/model"
	"instagram/common"
)

func (m *mySQLStorage) HasBeenReactionStory(ctx context.Context, data model.ReactionStory) (bool, error) {
	db := m.db.GetConnection()

	err := db.Table(model.ReactionStory{}.TableName()).Where("user_id = ? AND story_id = ?", data.UserId, data.StoryId).First(&data).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, model.ErrRecordReactStoryNotFound
		}
		return false, common.ErrDB(err)
	}
	return true, nil
}
