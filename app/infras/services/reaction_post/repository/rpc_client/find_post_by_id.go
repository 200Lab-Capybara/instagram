package rpc_client

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	reactionpostmodel "instagram/app/internals/services/reaction_post/model"
	"instagram/common"
)

type getPostRepo struct {
	db common.SQLDatabase
}

func NewGetPostRepo(db common.SQLDatabase) *getPostRepo {
	return &getPostRepo{db: db}
}

func (repo *getPostRepo) FindById(ctx context.Context, postId uuid.UUID) (*reactionpostmodel.Post, error) {
	var post reactionpostmodel.Post
	if err := repo.db.GetConnection().Table("posts").Where("id =?", postId).First(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, reactionpostmodel.ErrPostDoNotExist
		}

		return nil, common.ErrDB(err)
	}
	return &post, nil
}
