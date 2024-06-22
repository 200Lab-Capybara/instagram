package postsmysql

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	postsmodel "instagram/app/internals/services/posts/model"
	"instagram/common"
)

func (store *mysqlStorage) DecreaseLikeCount(ctx context.Context, postID uuid.UUID) (bool, error) {
	post := &postsmodel.Post{}
	err := store.db.GetConnection().WithContext(ctx).Table(post.TableName()).Where("id = ?", postID).Update("like_count", gorm.Expr("like_count - ?", 1)).Error

	if err != nil {
		return false, common.ErrDB(err)
	}

	return true, nil
}
