package common

import "gorm.io/gorm"

type SQLDatabase interface {
	GetConnection() *gorm.DB
}

type database struct {
	con *gorm.DB
}

func NewSQLDatabase(con *gorm.DB) SQLDatabase {
	return &database{
		con: con,
	}
}

func (d *database) GetConnection() *gorm.DB {
	return d.con.Session(&gorm.Session{NewDB: true})
}
