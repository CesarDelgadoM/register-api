package repository

import (
	"errors"

	"github.com/CesarDelgadoM/register-API/databases"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	connection *gorm.DB
}

func NewRepository[T any](T) *Repository[T] {
	return &Repository[T]{
		connection: databases.SelectDB("postgres").GetConnection(),
	}
}

func (repo *Repository[T]) SaveOne(model *T) (*T, error) {
	err := repo.connection.Debug().Create(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (repo *Repository[T]) SaveAll(model *[]T) (*[]T, error) {
	err := repo.connection.Debug().Create(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (repo *Repository[T]) GetAll() (*[]T, error) {
	var model *[]T
	err := repo.connection.Debug().Model(&model).Limit(100).Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (repo *Repository[T]) GetAllById(columnId string, id uint32) ([]T, error) {
	var model []T
	err := repo.connection.Debug().Model(&model).Where((columnId + " = ?"), id).Limit(100).Find(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (repo *Repository[T]) GetById(columnId string, id uint32) (*T, error) {
	var model *T
	err := repo.connection.Debug().Model(&model).Where((columnId + " = ?"), id).Take(&model).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (repo *Repository[T]) Update(columnId string, id uint32, columns map[string]interface{}) (*T, error) {
	var model *T
	err := repo.connection.Debug().Model(&model).Where((columnId + " = ?"), id).Updates(columns).Error
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (repo *Repository[T]) Delete(columnId string, id uint32) (int64, error) {
	var model *T
	db := repo.connection.Debug().Model(&model).Where((columnId + " = ?"), id).Delete(&model)
	if db.Error != nil {
		return 0, errors.New("not found")
	}
	return db.RowsAffected, nil
}
