package rpc_client

import (
	"context"
	"github.com/google/uuid"
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
	if err := repo.db.GetConnection().Table("stories").Where("id =?", sid).First(&story).Error; err != nil {
		return nil, err
	}
	return &story, nil
}
