package postsmysql

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	postsmodel "instagram/app/internals/services/posts/model"
	"instagram/common"
)

func (store *mysqlStorage) GetByID(ctx context.Context, postID uuid.UUID) (*postsmodel.Post, error) {
	post := postsmodel.Post{}
	err := store.db.GetConnection().WithContext(ctx).Where("id = ?", postID).First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, postsmodel.ErrPostNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &post, err
}
