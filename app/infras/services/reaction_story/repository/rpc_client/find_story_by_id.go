package rpc_client

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"instagram/app/internals/services/reaction_story/model"
	"instagram/common"
)

type getStoryRepo struct {
	db common.SQLDatabase
}

func NewGetStoryRepo(db common.SQLDatabase) *getStoryRepo {
	return &getStoryRepo{db: db}
}

func (repo *getStoryRepo) FindStoryById(ctx context.Context, sid uuid.UUID) (*model.Story, error) {
	var story model.Story
	if err := repo.db.GetConnection().Table(model.Story{}.TableName()).Where("id =?", sid).First(&story).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.ErrStoryNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &story, nil
}
