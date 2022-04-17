package database

import (
	"gorm.io/gorm"
)

type IDataBase interface {
	Connect()
	GetConnection() *gorm.DB
}

type DataBase struct {
	*gorm.DB
}

func (db *DataBase) GetConnection() *gorm.DB {
	return db.DB
}
