package rpc_client

import (
	"context"
	"github.com/google/uuid"
	"instagram/app/internals/services/reaction_post/model"
	"instagram/common"
)

type getPostRepo struct {
	db common.SQLDatabase
}

func NewGetPostRepo(db common.SQLDatabase) *getPostRepo {
	return &getPostRepo{db: db}
}

func (repo *getPostRepo) FindById(ctx context.Context, postId uuid.UUID) (*model.Post, error) {
	var post model.Post
	if err := repo.db.GetConnection().Table("posts").Where("id =?", postId).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}
