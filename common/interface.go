package common

import "gorm.io/gorm"

type SQLDatabase interface {
	GetConnection() *gorm.DB
}

type database struct {
	con *gorm.DB
}
