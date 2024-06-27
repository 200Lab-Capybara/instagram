package storymysql

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"instagram/app/internals/services/stories/model"
	"instagram/common"
)

func (storage *mysqlStorage) IncreaseLikeCount(ctx context.Context, storyId uuid.UUID) (bool, error)  {
	db := storage.db.GetConnection()
	err :=  db.Table(model.Story{}.TableName()).Where("id = ?",storyId).Update("like_count", gorm.Expr("like_count + ?", 1)).Error
	if err != nil {
		return false, common.ErrDB(err)
	}
	return true,nil
}