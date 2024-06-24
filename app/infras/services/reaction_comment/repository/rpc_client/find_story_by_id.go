package rpc_client

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	modelreactioncomment "instagram/app/internals/services/reaction_comment/model"
	"instagram/common"
)

type getCommentRepo struct {
	db common.SQLDatabase
}

func NewGetCommentRepo(db common.SQLDatabase) *getCommentRepo {
	return &getCommentRepo{db: db}
}

func (repo *getCommentRepo) FindCommentById(ctx context.Context, commentId uuid.UUID) (*modelreactioncomment.Comment, error) {
	var comment modelreactioncomment.Comment
	if err := repo.db.GetConnection().Table(modelreactioncomment.Comment{}.TableName()).
		Where("id = ?", commentId).First(&comment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, modelreactioncomment.ErrCommentNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &comment, nil
}
