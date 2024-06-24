package postsmysql

import "instagram/common"

type mysqlStorage struct {
	db common.SQLDatabase
}

func NewMysqlStorage(db common.SQLDatabase) *mysqlStorage {
	return &mysqlStorage{db: db}
}
