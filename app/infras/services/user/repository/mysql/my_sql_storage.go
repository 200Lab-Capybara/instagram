package usermysql

import (
	"github.com/nghiatrann0502/instagram-clone/common"
)

type mySQLStorage struct {
	db common.SQLDatabase
}

func NewMySQLStorage(db common.SQLDatabase) (*mySQLStorage, error) {

	return &mySQLStorage{
		db: db,
	}, nil
}
