package repository

import (
	"github.com/CesarDelgadoM/register-API/models"
)

type RegisterRepo struct {
	*Repository[models.Register]
}

func NewRegisterRepo() *RegisterRepo {
	return &RegisterRepo{
		NewRepository(models.Register{}),
	}
}

func (regrepo *RegisterRepo) GetByIdJoin(id uint32) (*models.Register, error) {
	var reg *models.Register
	err := regrepo.connection.Debug().Model(&reg).Select("*").Joins("JOIN \"objects\" ON registers.reg_id = objects.obj_reg_id").Where("reg_id = ?", id).Find(&reg).Error
	if err != nil {
		return nil, err
	}
	return reg, nil
}
