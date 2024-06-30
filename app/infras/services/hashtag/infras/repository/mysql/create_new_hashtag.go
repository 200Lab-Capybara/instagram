package hashtagsql

import (
	"context"
	hashtagmodel "instagram/app/internals/services/hashtag/model"
	"instagram/common"
)

//type mySQLStorage struct {
//	db common.SQLDatabase
//}
//
//func NewMySQLStorage(db common.SQLDatabase) *mySQLStorage {
//	return &mySQLStorage{
//		db: db,
//	}
//}

func (m *mySQLStorage) CreateHashTag(ctx context.Context, newHashTag hashtagmodel.Hashtag) (bool, error) {
	var data hashtagmodel.Hashtag
	db := m.db.GetConnection()
	if err := db.Table(data.TableName()).Create(&newHashTag).Error; err != nil {
		return false, common.ErrDB(err)
	}
	return true, nil
}
