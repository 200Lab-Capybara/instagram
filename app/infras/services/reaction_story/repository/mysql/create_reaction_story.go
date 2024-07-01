package reactstorymysql

import (
	"context"
	"github.com/google/uuid"
	"instagram/app/internals/services/reaction_story/model"
	"instagram/common"
	"time"
)

func (m *mySQLStorage) CreateNewReactionStory(ctx context.Context, sid uuid.UUID, uid uuid.UUID) (bool, error) {
	db := m.db.GetConnection()
	newRow := &model.ReactionStory{
		UserId:    uid,
		StoryId:   sid,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
	if err := db.Table(model.ReactionStory{}.TableName()).
		Create(newRow).Error; err != nil {
		return false, common.ErrDB(err)
	}
	return true, nil
}
