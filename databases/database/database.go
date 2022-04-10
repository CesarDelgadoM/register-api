package database

import "gorm.io/gorm"

type IDataBase interface {
	GetConnection() *gorm.DB
}

type DataBase struct {
	DB *gorm.DB
}

func (db *DataBase) GetConnection() *gorm.DB {
	return db.DB
}
