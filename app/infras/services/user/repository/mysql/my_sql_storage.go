package usermysql

import (
	"database/sql"
)

type mySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(db *sql.DB) (*mySQLStorage, error) {

	return &mySQLStorage{
		db: db,
	}, nil
}
