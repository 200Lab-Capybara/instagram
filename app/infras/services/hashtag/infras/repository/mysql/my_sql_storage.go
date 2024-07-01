package hashtagsql

import (
	"instagram/common"
)

type mySQLStorage struct {
	db common.SQLDatabase
}

func NewMySQLStorage(db common.SQLDatabase) *mySQLStorage {
	return &mySQLStorage{
		db: db,
	}
}
