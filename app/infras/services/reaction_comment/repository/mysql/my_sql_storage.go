package mysqlreactcomment

import "instagram/common"

// ket noi voi DB
type mySQLStorage struct {
	db common.SQLDatabase
}

func NewMySQLStorage(db common.SQLDatabase) *mySQLStorage {
	return &mySQLStorage{
		db: db,
	}
}
