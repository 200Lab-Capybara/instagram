package reactstorymysql

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"instagram/app/internals/services/reaction_story/model"
	"time"
)

func (m *mySQLStorage) HasBeenReactionStory(ctx context.Context, sid uuid.UUID, uid uuid.UUID) (*model.ReactionStory, error) {
	db := m.db.GetConnection()
	newRow := &model.ReactionStory{
		UserId:     uid,
		StoryId:    sid,
		Created_At: time.Now().UTC(),
		Updated_At: nil,
	}
	err := db.Table(model.ReactionStory{}.TableName()).Where("user_id = ? AND story_id = ?", uid, sid).First(newRow).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return newRow, nil
}
