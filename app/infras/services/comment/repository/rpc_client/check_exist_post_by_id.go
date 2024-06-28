package rpc_clientcomment

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	modelcomment "instagram/app/internals/services/comments/model"
	"instagram/common"
)

type getPostRepo struct {
	db common.SQLDatabase
}

func NewGetPostRepo(db common.SQLDatabase) *getPostRepo {
	return &getPostRepo{db: db}
}

func (repo *getPostRepo) CheckExistPostById(ctx context.Context, pid uuid.UUID) (bool, error) {
	var post modelcomment.Post

	if err := repo.db.GetConnection().Table(modelcomment.Post{}.TableName()).Where("id = ?", pid).First(&post).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, modelcomment.ErrPostNotFound
		}
		return false, common.ErrDB(err)
	}
	return true, nil
}
