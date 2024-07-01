package postsmysql

import (
	"context"
	"github.com/google/uuid"
	postsmodel "instagram/app/internals/services/posts/model"
	"instagram/common"
)

func (store *mysqlStorage) GetPostsByUserIds(ctx context.Context, ids []uuid.UUID, paging *common.Paging) ([]postsmodel.Post, error) {
	var data []postsmodel.Post
	db := store.db.GetConnection()

	db = db.Table(postsmodel.Post{}.TableName()).Where("status <> ?", postsmodel.PostDeleted).Where("user_id IN ?", ids)

	if err := db.Select("id").Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if v := paging.Cursor; v != "" {
		db = db.Where("created_at < ?", v)
	} else {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	}

	if err := db.Select("*").
		Order("created_at desc").
		Limit(paging.Limit).
		Find(&data).Error; err != nil {

		return nil, common.ErrDB(err)
	}

	if len(data) > 0 {
		paging.NextCursor = data[len(data)-1].CreatedAt.String()
	}

	return data, nil
}
