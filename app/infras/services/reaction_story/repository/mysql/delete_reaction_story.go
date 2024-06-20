package reactstorymysql

import (
	"context"
	"github.com/google/uuid"
	"instagram/app/internals/services/reaction_story/model"
)

func (m *mySQLStorage) DelReactionStory(ctx context.Context, sid uuid.UUID, uid uuid.UUID) (bool, error) {
	db := m.db.GetConnection()
	if err := db.Table(model.ReactionStory{}.TableName()).
		Where("story_id = ? AND user_id = ?", sid, uid).
		Delete(&model.ReactionStory{}).Error; err != nil {
		return false, err
	}

	return true, nil
}
