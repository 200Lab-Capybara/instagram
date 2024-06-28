package hashtagsql

import (
	"context"
	hashtagmodel "instagram/app/internals/services/hashtag/model"
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
	db := m.db.GetConnection()
	if err := db.Table(hashtagmodel.Hashtag{}.TableName()).Create(&newHashTag).Error; err != nil {
		return false, err
	}
	return true, nil
}
