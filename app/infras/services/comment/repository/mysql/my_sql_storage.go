package mysqlcomment

import (
	"instagram/common"
)

// ket noi voi database
type mysqlStorage struct {
	db common.SQLDatabase
}

func NewMySQLStorage(db common.SQLDatabase) *mysqlStorage {
	return &mysqlStorage{
		db: db,
	}
}
