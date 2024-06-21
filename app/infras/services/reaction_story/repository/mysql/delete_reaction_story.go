package reactstorymysql

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"instagram/app/internals/services/reaction_story/model"
	"instagram/common"
)

func (m *mySQLStorage) RemoveReactionStory(ctx context.Context, sid uuid.UUID, uid uuid.UUID) (bool, error) {
	db := m.db.GetConnection()
	if err := db.Table(model.ReactionStory{}.TableName()).
		Where("story_id = ? AND user_id = ?", sid, uid).
		Delete(&model.ReactionStory{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, model.ErrRecordReactStoryNotFound
		}
		return false, common.ErrDB(err)
	}

	return true, nil
}
