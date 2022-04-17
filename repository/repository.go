package repository

import (
	"github.com/CesarDelgadoM/register-API/databases"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	connection *gorm.DB
}

func New[T any](T) *Repository[T] {
	return &Repository[T]{
		connection: databases.SelectDB("postgres").GetConnection(),
	}
}

func (repo *Repository[T]) Save(model *T) (*T, error) {
	err := repo.connection.Debug().Create(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (repo *Repository[T]) GetAll() (*[]T, error) {
	var model *[]T
	err := repo.connection.Debug().Model(&model).Limit(100).Find(model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (repo *Repository[T]) GetById(nameColumn string, id uint32) (*T, error) {
	var model *T
	err := repo.connection.Debug().Model(&model).Where((nameColumn + " = ?"), id).Take(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (repo *Repository[T]) Update(columnId string, id uint32, model *T, columns map[string]interface{}) (*T, error) {
	err := repo.connection.Debug().Model(&model).Where((columnId + " = ?"), id).Updates(columns).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (repo *Repository[T]) Delete(ColumnId string, id uint32) {

}
