package reactstorymysql

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"instagram/app/internals/services/reaction_story/model"
	"instagram/common"
	"time"
)

func (m *mySQLStorage) HasBeenReactionStory(ctx context.Context, sid uuid.UUID, uid uuid.UUID) (bool, error) {
	db := m.db.GetConnection()
	newRow := &model.ReactionStory{
		UserId:    uid,
		StoryId:   sid,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: nil,
	}
	err := db.Table(model.ReactionStory{}.TableName()).Where("user_id = ? AND story_id = ?", uid, sid).First(newRow).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, model.ErrRecordReactStoryNotFound
		}
		return false, common.ErrDB(err)
	}
	return true, nil
}
