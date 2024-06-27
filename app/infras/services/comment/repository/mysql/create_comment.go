package mysqlcomment

import (
	"context"
	"github.com/google/uuid"
	modelcomment "instagram/app/internals/services/comments/model"
	"instagram/common"
)

func (content *mysqlStorage) CreateComment(ctx context.Context, dto modelcomment.Comment) (*uuid.UUID, error) {
	//tao phuong thuc ket noi den db
	db := content.db.GetConnection()

	if err := db.Table(dto.TableName()).Create(dto).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &dto.Id, nil
}
