package postsmysql

import (
	"context"
	"github.com/google/uuid"
	postsmodel "instagram/app/internals/services/posts/model"
	"instagram/common"
)

func (s *mysqlStorage) CreatePost(ctx context.Context, post *postsmodel.Post) (*uuid.UUID, error) {

	db := s.db.GetConnection()
	if err := db.Table(post.TableName()).Create(post).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return &post.ID, nil
}
