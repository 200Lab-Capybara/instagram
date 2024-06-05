package usermysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type mySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(db *sql.DB) (*mySQLStorage, error) {

	db, err := sql.Open("mysql", "capybara:my_secret@tcp(localhost:3306)/users?parseTime=true")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &mySQLStorage{
		db: db,
	}, nil
}
