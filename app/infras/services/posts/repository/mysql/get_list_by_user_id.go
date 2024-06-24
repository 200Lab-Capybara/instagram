package postsmysql

import (
	"context"
	"github.com/google/uuid"
	postsmodel "instagram/app/internals/services/posts/model"
	"instagram/common"
)

func (store *mysqlStorage) GetListPostByUserId(ctx context.Context, userId uuid.UUID, paging *common.Paging) ([]postsmodel.Post, error) {
	var data []postsmodel.Post
	db := store.db.GetConnection()

	if err := db.Select("id").Where("user_id=?", userId).Table(postsmodel.Post{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if v := paging.FakeCursor; v != "" {
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
