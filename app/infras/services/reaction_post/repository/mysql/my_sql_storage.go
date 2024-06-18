package reactionpostmysql

import (
	"github.com/nghiatrann0502/instagram-clone/common"
)

type mySQLStorage struct {
	db common.SQLDatabase
}

func NewMySQLStorage(db common.SQLDatabase) *mySQLStorage {
	return &mySQLStorage{
		db: db,
	}
}
