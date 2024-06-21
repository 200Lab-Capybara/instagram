package common

import "gorm.io/gorm"

func NewSQLDatabase(con *gorm.DB) SQLDatabase {
	return &database{
		con: con,
	}
}

func (d *database) GetConnection() *gorm.DB {
	return d.con.Session(&gorm.Session{NewDB: true})
}
