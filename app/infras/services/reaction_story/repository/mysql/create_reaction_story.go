package reactstorymysql

import (
	"context"
	"github.com/google/uuid"
	"instagram/app/internals/services/reaction_story/model"
	"time"
)

func (m *mySQLStorage) CreateNewReactionStory(ctx context.Context, sid uuid.UUID, uid uuid.UUID) (bool, error) {
	db := m.db.GetConnection()
	newRow := model.ReactionStory{
		UserId:     uid,
		StoryId:    sid,
		Created_At: time.Now().UTC(),
		Updated_At: nil,
	}
	if err := db.Table(model.ReactionStory{}.TableName()).
		Create(newRow).Error; err != nil {
		return false, err
	}
	return true, nil
}
